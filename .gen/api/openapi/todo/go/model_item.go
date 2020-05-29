/*
 * Todo API
 *
 * The Todo API manages a list of todo items as described by the TodoMVC backend project: http://todobackend.com 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

type Item struct {

	Id string `json:"id"`

	Title string `json:"title"`

	Completed bool `json:"completed"`

	Order int32 `json:"order"`

	Url string `json:"url"`
}