//go:generate go run ../../generate/tags/main.go -AWSSDKVersion=2 -ListTagsInIDElem=ResourceName -ServiceTagsSlice -TagInIDElem=ResourceName -UpdateTags
//go:generate go run ../../generate/servicepackage/main.go
// ONLY generate directives and package declaration! Do not add anything else to this file.

package lightsail
