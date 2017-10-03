///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
package imagemanager

import (
	"fmt"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	log "github.com/sirupsen/logrus"

	entitystore "gitlab.eng.vmware.com/serverless/serverless/pkg/entity-store"
	"gitlab.eng.vmware.com/serverless/serverless/pkg/image-manager/gen/models"
	"gitlab.eng.vmware.com/serverless/serverless/pkg/image-manager/gen/restapi/operations"
	baseimage "gitlab.eng.vmware.com/serverless/serverless/pkg/image-manager/gen/restapi/operations/base_image"
	"gitlab.eng.vmware.com/serverless/serverless/pkg/image-manager/gen/restapi/operations/image"
	"gitlab.eng.vmware.com/serverless/serverless/pkg/trace"
)

// ImageManagerFlags are configuration flags for the image manager
var ImageManagerFlags = struct {
	DbFile       string `long:"db-file" description:"Path to BoltDB file" default:"./db.bolt"`
	OrgID        string `long:"organization" description:"(temporary) Static organization id" default:"serverless"`
	K8sConfig    string `long:"kubeconfig" description:"Path to kubernetes config file"`
	K8sNamespace string `long:"namespace" description:"Kubernetes namespace for jobs" default:"default"`
}{}

var statusMap = map[models.Status]entitystore.Status{
	models.StatusCREATING:    StatusCREATING,
	models.StatusDELETED:     StatusDELETED,
	models.StatusERROR:       StatusERROR,
	models.StatusINITIALIZED: StatusINITIALIZED,
	models.StatusREADY:       StatusREADY,
}

var reverseStatusMap = make(map[entitystore.Status]models.Status)

func initializeStatusMap() {
	defer trace.Trace("initializeStatusMap")()
	for k, v := range statusMap {
		reverseStatusMap[v] = k
	}
}

func baseImageEntityToModel(e *BaseImage) *models.BaseImage {
	defer trace.Trace("baseImageEntityToModel")()
	var tags []*models.Tag
	for k, v := range e.Tags {
		tags = append(tags, &models.Tag{Key: k, Value: v})
	}

	m := models.BaseImage{
		CreatedTime: e.CreatedTime.Unix(),
		DockerURL:   swag.String(e.DockerURL),
		ID:          strfmt.UUID(e.ID),
		Public:      swag.Bool(e.Public),
		Name:        swag.String(e.Name),
		Status:      reverseStatusMap[e.Status],
		Tags:        tags,
		Reason:      e.Reason,
	}
	return &m
}

func baseImageModelToEntity(m *models.BaseImage) *BaseImage {
	defer trace.Trace("baseImageModelToEntity")()
	tags := make(map[string]string)
	for _, t := range m.Tags {
		tags[t.Key] = t.Value
	}
	e := BaseImage{
		BaseEntity: entitystore.BaseEntity{
			OrganizationID: ImageManagerFlags.OrgID,
			Name:           *m.Name,
			Tags:           tags,
			Status:         statusMap[m.Status],
			Reason:         m.Reason,
		},
		DockerURL: *m.DockerURL,
		Public:    *m.Public,
	}
	return &e
}

func imageEntityToModel(e *Image) *models.Image {
	defer trace.Trace("imageEntityToModel")()
	var tags []*models.Tag
	for k, v := range e.Tags {
		tags = append(tags, &models.Tag{Key: k, Value: v})
	}

	m := models.Image{
		CreatedTime:   e.CreatedTime.Unix(),
		BaseImageName: swag.String(e.BaseImageName),
		DockerURL:     e.DockerURL,
		ID:            strfmt.UUID(e.ID),
		Name:          swag.String(e.Name),
		Status:        reverseStatusMap[e.Status],
		Tags:          tags,
		Reason:        e.Reason,
	}
	return &m
}

func imageModelToEntity(m *models.Image) *Image {
	defer trace.Trace("imageModelToEntity")()
	tags := make(map[string]string)
	for _, t := range m.Tags {
		tags[t.Key] = t.Value
	}
	e := Image{
		BaseEntity: entitystore.BaseEntity{
			OrganizationID: ImageManagerFlags.OrgID,
			Name:           *m.Name,
			Tags:           tags,
			Status:         statusMap[m.Status],
			Reason:         m.Reason,
		},
		DockerURL:     m.DockerURL,
		BaseImageName: *m.BaseImageName,
	}
	return &e
}

// Handlers encapsulates the image manager handlers
type Handlers struct {
	baseImageBuilder *BaseImageBuilder
}

// NewHandlers is the constructor for the Handlers type
func NewHandlers(baseImageBuilder *BaseImageBuilder) *Handlers {
	defer trace.Trace("NewHandlers")()
	return &Handlers{
		baseImageBuilder: baseImageBuilder,
	}
}

// ConfigureHandlers registers the image manager handlers to the API
func (h *Handlers) ConfigureHandlers(api middleware.RoutableAPI, store entitystore.EntityStore) {
	defer trace.Trace("ConfigureHandlers")()
	a, ok := api.(*operations.ImageManagerAPI)
	if !ok {
		panic("Cannot configure api")
	}

	initializeStatusMap()

	a.BaseImageAddBaseImageHandler = baseimage.AddBaseImageHandlerFunc(func(params baseimage.AddBaseImageParams) middleware.Responder {
		defer trace.Trace("BaseImageAddBaseImageHandler")()
		baseImageRequest := params.Body
		e := baseImageModelToEntity(baseImageRequest)
		e.Status = StatusINITIALIZED
		_, err := store.Add(e)
		if err != nil {
			log.Errorf("Store error when adding base image: %+v", err)
			return baseimage.NewAddBaseImageBadRequest()
		}
		if h.baseImageBuilder != nil {
			h.baseImageBuilder.baseImageChannel <- *e
		}
		m := baseImageEntityToModel(e)
		return baseimage.NewAddBaseImageCreated().WithPayload(m)
	})

	a.BaseImageGetBaseImageByNameHandler = baseimage.GetBaseImageByNameHandlerFunc(func(params baseimage.GetBaseImageByNameParams) middleware.Responder {
		defer trace.Trace("BaseImageGetBaseImageByNameHandler")()
		e := BaseImage{}
		err := store.Get(ImageManagerFlags.OrgID, params.BaseImageName, &e)
		if err != nil {
			log.Warnf("Received GET for non-existent base image %s", params.BaseImageName)
			log.Debugf("Error returned by store.Get: %+v", err)
			return baseimage.NewGetBaseImageByNameNotFound()
		}
		m := baseImageEntityToModel(&e)
		return baseimage.NewGetBaseImageByNameOK().WithPayload(m)
	})

	a.BaseImageGetBaseImagesHandler = baseimage.GetBaseImagesHandlerFunc(func(params baseimage.GetBaseImagesParams) middleware.Responder {
		defer trace.Trace("BaseImageGetBaseImagesHandler")()
		var images []BaseImage
		err := store.List(ImageManagerFlags.OrgID, nil, &images)
		if err != nil {
			log.Errorf("Store error when listing base images: %+v", err)
			return baseimage.NewGetBaseImagesDefault(500)
		}
		var imageModels []*models.BaseImage
		for _, image := range images {
			imageModels = append(imageModels, baseImageEntityToModel(&image))
		}
		return baseimage.NewGetBaseImagesOK().WithPayload(imageModels)
	})

	a.ImageAddImageHandler = image.AddImageHandlerFunc(func(params image.AddImageParams) middleware.Responder {
		defer trace.Trace("ImageAddImageHandler")()
		imageRequest := params.Body

		bi := BaseImage{}
		err := store.Get(ImageManagerFlags.OrgID, *imageRequest.BaseImageName, &bi)
		if err != nil {
			log.Warnf("Unable to add image, base image %s does not exist", *imageRequest.BaseImageName)
			log.Debugf("Error returned by store.Get: %+v", err)
			return image.NewAddImageBadRequest().WithPayload(&models.Error{Message: swag.String("Base image missing")})
		}
		if bi.Status != StatusREADY {
			log.Debugf("Base image %s not in ready status, actual status: %s", bi.Name, bi.Status)
			return image.NewAddImageBadRequest().WithPayload(&models.Error{Message: swag.String(fmt.Sprintf("Base image must be status %v", StatusREADY))})
		}

		e := imageModelToEntity(imageRequest)
		e.Status = StatusREADY
		e.DockerURL = bi.DockerURL
		_, err = store.Add(e)
		if err != nil {
			log.Errorf("Store error while adding new image: %+v", err)
			return image.NewAddImageBadRequest().WithPayload(&models.Error{Message: swag.String(err.Error())})
		}
		m := imageEntityToModel(e)
		return image.NewAddImageCreated().WithPayload(m)
	})

	a.ImageGetImageByNameHandler = image.GetImageByNameHandlerFunc(func(params image.GetImageByNameParams) middleware.Responder {
		defer trace.Trace("ImageGetImageByNameHandler")()
		e := Image{}
		err := store.Get(ImageManagerFlags.OrgID, params.ImageName, &e)
		if err != nil {
			log.Warnf("Received GET for non-existentimage %s", params.ImageName)
			log.Debugf("Error returned by store.Get: %+v", err)
			return image.NewGetImageByNameNotFound()
		}
		m := imageEntityToModel(&e)
		return image.NewGetImageByNameOK().WithPayload(m)
	})

	a.ImageGetImagesHandler = image.GetImagesHandlerFunc(func(params image.GetImagesParams) middleware.Responder {
		defer trace.Trace("ImageGetImagesHandler")()
		var images []Image
		err := store.List(ImageManagerFlags.OrgID, nil, &images)
		if err != nil {
			log.Errorf("Store error when listing images: %+v", err)
			return image.NewGetImagesDefault(500).WithPayload(&models.Error{Message: swag.String(err.Error())})
		}
		var imageModels []*models.Image
		for _, image := range images {
			imageModels = append(imageModels, imageEntityToModel(&image))
		}
		return image.NewGetImagesOK().WithPayload(imageModels)
	})

	a.ServerShutdown = func() {
		defer trace.Trace("ServerShutdown")()
		h.baseImageBuilder.done <- true
	}
}