// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `swagger: "2.0"
info:
  title: Your API
  description: API documentation for the Post and Category resources
  version: 1.0.0
host: localhost:8080
basePath: /api/v1
schemes:
  - http
paths:
  /posts:
    get:
      summary: Get all posts
      description: Retrieves a list of all posts
      produces:
        - application/json
      responses:
        200:
          description: Successful operation
          schema:
            type: array
            items:
              $ref: "#/definitions/Post"
    post:
      summary: Create a new post
      description: Creates a new post with the provided data
      consumes:
        - application/json
      produces:
        - application/json
      parameters:
        - name: post
          in: body
          description: Post object
          required: true
          schema:
            $ref: "#/definitions/Post"
      responses:
        201:
          description: Post created successfully
          schema:
            $ref: "#/definitions/Post"
  /posts/{id}:
    get:
      summary: Get a post by ID
      description: Retrieves a post by its ID
      produces:
        - application/json
      parameters:
        - name: id
          in: path
          description: ID of the post to retrieve
          required: true
          type: integer
          format: int64
      responses:
        200:
          description: Successful operation
          schema:
            $ref: "#/definitions/Post"
    put:
      summary: Update a post by ID
      description: Updates a post with the provided data
      consumes:
        - application/json
      produces:
        - application/json
      parameters:
        - name: id
          in: path
          description: ID of the post to update
          required: true
          type: integer
          format: int64
        - name: post
          in: body
          description: Post object
          required: true
          schema:
            $ref: "#/definitions/Post"
      responses:
        200:
          description: Post updated successfully
          schema:
            $ref: "#/definitions/Post"
    delete:
      summary: Delete a post by ID
      description: Deletes a post by its ID
      parameters:
        - name: id
          in: path
          description: ID of the post to delete
          required: true
          type: integer
          format: int64
      responses:
        204:
          description: Post deleted successfully
  /category:
    get:
      summary: Get all categories
      description: Retrieves a list of all categories
      produces:
        - application/json
      responses:
        200:
          description: Successful operation
          schema:
            type: array
            items:
              $ref: "#/definitions/Category"
    post:
      summary: Create a new category
      description: Creates a new category with the provided data
      consumes:
        - application/json
      produces:
        - application/json
      parameters:
        - name: category
          in: body
          description: Category object
          required: true
          schema:
            $ref: "#/definitions/Category"
      responses:
        201:
          description: Category created successfully
          schema:
            $ref: "#/definitions/Category"
  /category/{id}:
    get:
      summary: Get a category by ID
      description: Retrieves a category by its ID
      produces:
        - application/json
      parameters:
        - name: id
          in: path
          description: ID of the category to retrieve
          required: true
          type: integer
          format: int64
      responses:
        200:
          description: Successful operation
          schema:
            $ref: "#/definitions/Category"
    put:
      summary: Update a category by ID
      description: Updates a category with the provided data
      consumes:
        - application/json
      produces:
        - application/json
      parameters:
        - name: id
          in: path
          description: ID of the category to update
          required: true
          type: integer
          format: int64
        - name: category
          in: body
          description: Category object
          required: true
          schema:
            $ref: "#/definitions/Category"
      responses:
        200:
          description: Category updated successfully
          schema:
            $ref: "#/definitions/Category"
    delete:
      summary: Delete a category by ID
      description: Deletes a category by its ID
      parameters:
        - name: id
          in: path
          description: ID of the category to delete
          required: true
          type: integer
          format: int64
      responses:
        204:
          description: Category deleted successfully
definitions:
  Post:
    type: object
    properties:
      id:
        type: integer
        format: int64
      title:
        type: string
      text:
        type: string
      create_at:
        type: string
        format: date-time
      modified_at:
        type: string
        format: date-time
      categories:
        type: array
        items:
          $ref: "#/definitions/Category"
  Category:
    type: object
    properties:
      id:
        type: integer
        format: int64
      name:
        type: string
      create_at:
        type: string
        format: date-time
      updated_at:
        type: string
        format: date-time
      posts:
        type: array
        items:
          $ref: "#/definitions/Post"
`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v1.0.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "RESFUL API ",
	Description:      "This documentation for Echo REST blog.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
