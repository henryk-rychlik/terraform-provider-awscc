// Code generated by generators/resource/main.go; DO NOT EDIT.

package appmesh

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/generic"
	"github.com/hashicorp/terraform-provider-aws-cloudapi/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("aws_appmesh_virtual_service", awsAppmeshVirtualService)
}

// awsAppmeshVirtualService returns the Terraform aws_appmesh_virtual_service resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::AppMesh::VirtualService resource type.
func awsAppmeshVirtualService(ctx context.Context) (tfsdk.ResourceType, error) {
	// Subproperty definitions.

	// Definition: Tag
	// Property: Key
	// CloudFormation resource type schema:
	/*
	   {
	     "type": "string"
	   }
	*/
	attr2859957898 := schema.Attribute{}
	attr2859957898.Type = types.StringType
	attr2859957898.Optional = true

	// Definition: Tag
	// Property: Value
	// CloudFormation resource type schema:
	/*
	   {
	     "type": "string"
	   }
	*/
	attr3708964976 := schema.Attribute{}
	attr3708964976.Type = types.StringType
	attr3708964976.Optional = true

	// Property references for Tag:
	attr4169356339 := make(map[string]schema.Attribute, 2)
	attr4169356339["key"] = attr2859957898
	attr4169356339["value"] = attr3708964976

	// Definition: VirtualNodeServiceProvider
	// Property: VirtualNodeName
	// CloudFormation resource type schema:
	/*
	   {
	     "type": "string"
	   }
	*/
	attr1959102114 := schema.Attribute{}
	attr1959102114.Type = types.StringType
	attr1959102114.Optional = true

	// Property references for VirtualNodeServiceProvider:
	attr2722040410 := make(map[string]schema.Attribute, 1)
	attr2722040410["virtual_node_name"] = attr1959102114

	// Definition: VirtualRouterServiceProvider
	// Property: VirtualRouterName
	// CloudFormation resource type schema:
	/*
	   {
	     "type": "string"
	   }
	*/
	attr1625712818 := schema.Attribute{}
	attr1625712818.Type = types.StringType
	attr1625712818.Optional = true

	// Property references for VirtualRouterServiceProvider:
	attr1187211989 := make(map[string]schema.Attribute, 1)
	attr1187211989["virtual_router_name"] = attr1625712818

	// Definition: VirtualServiceProvider
	// Property: VirtualNode
	// CloudFormation resource type schema:
	/*
	   {
	     "$ref": "#/definitions/VirtualNodeServiceProvider"
	   }
	*/
	attr956383445 := schema.Attribute{}
	attr956383445.Attributes = schema.SingleNestedAttributes(attr2722040410)
	attr956383445.Optional = true

	// Definition: VirtualServiceProvider
	// Property: VirtualRouter
	// CloudFormation resource type schema:
	/*
	   {
	     "$ref": "#/definitions/VirtualRouterServiceProvider"
	   }
	*/
	attr3217696956 := schema.Attribute{}
	attr3217696956.Attributes = schema.SingleNestedAttributes(attr1187211989)
	attr3217696956.Optional = true

	// Property references for VirtualServiceProvider:
	attr2814324958 := make(map[string]schema.Attribute, 2)
	attr2814324958["virtual_node"] = attr956383445
	attr2814324958["virtual_router"] = attr3217696956

	// Definition: VirtualServiceSpec
	// Property: Provider
	// CloudFormation resource type schema:
	/*
	   {
	     "$ref": "#/definitions/VirtualServiceProvider"
	   }
	*/
	attr2865320949 := schema.Attribute{}
	attr2865320949.Attributes = schema.SingleNestedAttributes(attr2814324958)
	attr2865320949.Optional = true

	// Property references for VirtualServiceSpec:
	attr512434520 := make(map[string]schema.Attribute, 1)
	attr512434520["provider"] = attr2865320949

	// Root property definition.

	// Definition: (Root)
	// Property: Arn
	// CloudFormation resource type schema:
	/*
	   {
	     "type": "string"
	   }
	*/
	attr3018501457 := schema.Attribute{}
	attr3018501457.Type = types.StringType
	attr3018501457.Computed = true

	// Definition: (Root)
	// Property: Id
	// PrimaryIdentifier: true
	// CloudFormation resource type schema:
	/*
	   {
	     "type": "string"
	   }
	*/
	attr2710565547 := schema.Attribute{}
	attr2710565547.Type = types.StringType
	attr2710565547.Computed = true

	// Definition: (Root)
	// Property: MeshName
	// CloudFormation resource type schema:
	/*
	   {
	     "type": "string"
	   }
	*/
	attr3520476156 := schema.Attribute{}
	attr3520476156.Type = types.StringType
	attr3520476156.Required = true

	// Definition: (Root)
	// Property: MeshOwner
	// CloudFormation resource type schema:
	/*
	   {
	     "type": "string"
	   }
	*/
	attr1912911946 := schema.Attribute{}
	attr1912911946.Type = types.StringType
	attr1912911946.Optional = true

	// Definition: (Root)
	// Property: ResourceOwner
	// CloudFormation resource type schema:
	/*
	   {
	     "type": "string"
	   }
	*/
	attr4099733265 := schema.Attribute{}
	attr4099733265.Type = types.StringType
	attr4099733265.Computed = true

	// Definition: (Root)
	// Property: Spec
	// CloudFormation resource type schema:
	/*
	   {
	     "$ref": "#/definitions/VirtualServiceSpec"
	   }
	*/
	attr865609681 := schema.Attribute{}
	attr865609681.Attributes = schema.SingleNestedAttributes(attr512434520)
	attr865609681.Required = true

	// Definition: (Root)
	// Property: Tags
	// CloudFormation resource type schema:
	/*
	   {
	     "items": {
	       "$ref": "#/definitions/Tag"
	     },
	     "type": "array",
	     "uniqueItems": false
	   }
	*/
	attr2151425999 := schema.Attribute{}
	attr2151425999Options := schema.ListNestedAttributesOptions{}
	attr2151425999.Attributes = schema.ListNestedAttributes(attr4169356339, attr2151425999Options)
	attr2151425999.Optional = true

	// Definition: (Root)
	// Property: Uid
	// CloudFormation resource type schema:
	/*
	   {
	     "type": "string"
	   }
	*/
	attr1173502620 := schema.Attribute{}
	attr1173502620.Type = types.StringType
	attr1173502620.Computed = true

	// Definition: (Root)
	// Property: VirtualServiceName
	// CloudFormation resource type schema:
	/*
	   {
	     "type": "string"
	   }
	*/
	attr1309939279 := schema.Attribute{}
	attr1309939279.Type = types.StringType
	attr1309939279.Required = true

	// Property references for (Root):
	attr1863345718 := make(map[string]schema.Attribute, 9)
	attr1863345718["arn"] = attr3018501457
	attr1863345718["id"] = attr2710565547
	attr1863345718["mesh_name"] = attr3520476156
	attr1863345718["mesh_owner"] = attr1912911946
	attr1863345718["resource_owner"] = attr4099733265
	attr1863345718["spec"] = attr865609681
	attr1863345718["tags"] = attr2151425999
	attr1863345718["uid"] = attr1173502620
	attr1863345718["virtual_service_name"] = attr1309939279

	// Resource instance unique identifier.
	attr1863345718["identifier"] = schema.Attribute{
		Type:        types.StringType,
		Computed:    true,
		Description: "The resource instance's unique identifier.",
	}

	schema := schema.Schema{
		Version:    1,
		Attributes: attr1863345718,
	}

	resourceType := generic.NewResourceType(
		"AWS::AppMesh::VirtualService",
		"aws_appmesh_virtual_service",
		schema,
	)

	return resourceType, nil
}
