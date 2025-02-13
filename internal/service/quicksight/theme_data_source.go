package quicksight

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/quicksight"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/create"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKDataSource("aws_quicksight_theme", name="Theme")
func DataSourceTheme() *schema.Resource {
	return &schema.Resource{
		ReadWithoutTimeout: dataSourceThemeRead,

		SchemaFunc: func() map[string]*schema.Schema {
			return map[string]*schema.Schema{
				"arn": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"aws_account_id": {
					Type:         schema.TypeString,
					Optional:     true,
					Computed:     true,
					ValidateFunc: verify.ValidAccountID,
				},
				"base_theme_id": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"configuration": { // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_ThemeConfiguration.html
					Type:     schema.TypeList,
					Computed: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"data_color_palette": { // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DataColorPalette.html
								Type:     schema.TypeList,
								Computed: true,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"colors": {
											Type:     schema.TypeList,
											Computed: true,
											Elem: &schema.Schema{
												Type: schema.TypeString,
											},
										},
										"empty_fill_color": {
											Type:     schema.TypeString,
											Computed: true,
										},
										"min_max_gradient": {
											Type:     schema.TypeList,
											Computed: true,
											Elem: &schema.Schema{
												Type: schema.TypeString,
											},
										},
									},
								},
							},
							"sheet": { // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_SheetStyle.html
								Type:     schema.TypeList,
								Computed: true,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"tile": { // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_TileStyle.html
											Type:     schema.TypeList,
											Computed: true,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"border": { // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_BorderStyle.html
														Type:     schema.TypeList,
														Computed: true,
														Elem: &schema.Resource{
															Schema: map[string]*schema.Schema{
																"show": {
																	Type:     schema.TypeBool,
																	Computed: true,
																},
															},
														},
													},
												},
											},
										},
										"tile_layout": { // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_TileLayoutStyle.html
											Type:     schema.TypeList,
											Computed: true,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"gutter": { // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_GutterStyle.html
														Type:     schema.TypeList,
														Computed: true,
														Elem: &schema.Resource{
															Schema: map[string]*schema.Schema{
																"show": {
																	Type:     schema.TypeBool,
																	Computed: true,
																},
															},
														},
													},
													"margin": { // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_MarginStyle.html
														Type:     schema.TypeList,
														Computed: true,
														Elem: &schema.Resource{
															Schema: map[string]*schema.Schema{
																"show": {
																	Type:     schema.TypeBool,
																	Computed: true,
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
							"typography": { // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_Typography.html
								Type:     schema.TypeList,
								Computed: true,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"font_families": { // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_Font.html
											Type:     schema.TypeList,
											Computed: true,
											Elem: &schema.Resource{
												Schema: map[string]*schema.Schema{
													"font_family": {
														Type:     schema.TypeString,
														Computed: true,
													},
												},
											},
										},
									},
								},
							},
							"ui_color_palette": { // https://docs.aws.amazon.com/quicksight/latest/APIReference/API_UIColorPalette.html
								Type:     schema.TypeList,
								Computed: true,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"accent": {
											Type:     schema.TypeString,
											Computed: true,
										},
										"accent_foreground": {
											Type:     schema.TypeString,
											Computed: true,
										},
										"danger": {
											Type:     schema.TypeString,
											Computed: true,
										},
										"danger_foreground": {
											Type:     schema.TypeString,
											Computed: true,
										},
										"dimension": {
											Type:     schema.TypeString,
											Computed: true,
										},
										"dimension_foreground": {
											Type:     schema.TypeString,
											Computed: true,
										},
										"measure": {
											Type:     schema.TypeString,
											Computed: true,
										},
										"measure_foreground": {
											Type:     schema.TypeString,
											Computed: true,
										},
										"primary_background": {
											Type:     schema.TypeString,
											Computed: true,
										},
										"primary_foreground": {
											Type:     schema.TypeString,
											Computed: true,
										},
										"secondary_background": {
											Type:     schema.TypeString,
											Computed: true,
										},
										"secondary_foreground": {
											Type:     schema.TypeString,
											Computed: true,
										},
										"success": {
											Type:     schema.TypeString,
											Computed: true,
										},
										"success_foreground": {
											Type:     schema.TypeString,
											Computed: true,
										},
										"warning": {
											Type:     schema.TypeString,
											Computed: true,
										},
										"warning_foreground": {
											Type:     schema.TypeString,
											Computed: true,
										},
									},
								},
							},
						},
					},
				},
				"created_time": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"theme_id": {
					Type:     schema.TypeString,
					Required: true,
				},
				"last_updated_time": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"name": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"permissions": {
					Type:     schema.TypeList,
					Computed: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"actions": {
								Type:     schema.TypeSet,
								Computed: true,
								Elem:     &schema.Schema{Type: schema.TypeString},
							},
							"principal": {
								Type:     schema.TypeString,
								Computed: true,
							},
						},
					},
				},
				"status": {
					Type:     schema.TypeString,
					Computed: true,
				},
				names.AttrTags: tftags.TagsSchemaComputed(),
				"version_description": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"version_number": {
					Type:     schema.TypeInt,
					Computed: true,
				},
			}
		},
	}
}

const (
	DSNameTheme = "Theme Data Source"
)

func dataSourceThemeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).QuickSightConn(ctx)

	awsAccountId := meta.(*conns.AWSClient).AccountID
	if v, ok := d.GetOk("aws_account_id"); ok {
		awsAccountId = v.(string)
	}
	themeId := d.Get("theme_id").(string)

	id := createThemeId(awsAccountId, themeId)

	out, err := FindThemeByID(ctx, conn, id)

	if err != nil {
		return create.DiagError(names.QuickSight, create.ErrActionReading, ResNameTheme, d.Id(), err)
	}

	d.SetId(id)
	d.Set("arn", out.Arn)
	d.Set("aws_account_id", awsAccountId)
	d.Set("base_theme_id", out.Version.BaseThemeId)
	d.Set("created_time", out.CreatedTime.Format(time.RFC3339))
	d.Set("last_updated_time", out.LastUpdatedTime.Format(time.RFC3339))
	d.Set("name", out.Name)
	d.Set("status", out.Version.Status)
	d.Set("theme_id", out.ThemeId)
	d.Set("version_description", out.Version.Description)
	d.Set("version_number", out.Version.VersionNumber)

	if err := d.Set("configuration", flattenThemeConfiguration(out.Version.Configuration)); err != nil {
		return diag.Errorf("setting configuration: %s", err)
	}

	permsResp, err := conn.DescribeThemePermissionsWithContext(ctx, &quicksight.DescribeThemePermissionsInput{
		AwsAccountId: aws.String(awsAccountId),
		ThemeId:      aws.String(themeId),
	})

	if err != nil {
		return diag.Errorf("describing QuickSight Theme (%s) Permissions: %s", d.Id(), err)
	}

	if err := d.Set("permissions", flattenPermissions(permsResp.Permissions)); err != nil {
		return diag.Errorf("setting permissions: %s", err)
	}

	return nil
}
