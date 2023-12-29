/*
NuoDB Control Plane REST API

NuoDB Control Plane (CP) allows users to create and manage NuoDB databases remotely using a Database as a Service (DBaaS) model.

API version: 2.3.0
Contact: NuoDB.Support@3ds.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nuodbaas

import (
	"encoding/json"
)

// checks if the DatabaseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseModel{}

// DatabaseModel struct for DatabaseModel
type DatabaseModel struct {
	Organization *string `json:"organization,omitempty"`
	Project *string `json:"project,omitempty"`
	Name *string `json:"name,omitempty"`
	// User-defined labels attached to the resource that can be used for filtering
	Labels *map[string]string `json:"labels,omitempty"`
	// The service tier for the database. If omitted, the project service tier is inherited.
	Tier *string `json:"tier,omitempty"`
	Maintenance *MaintenanceModel `json:"maintenance,omitempty"`
	Properties *DatabasePropertiesModel `json:"properties,omitempty"`
	// The version of the resource. When specified in a `PUT` request payload, indicates that the resoure should be updated, and is used by the system to guard against concurrent updates.
	ResourceVersion *string `json:"resourceVersion,omitempty"`
	Status *DatabaseStatusModel `json:"status,omitempty"`
}

// NewDatabaseModel instantiates a new DatabaseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseModel() *DatabaseModel {
	this := DatabaseModel{}
	return &this
}

// NewDatabaseModelWithDefaults instantiates a new DatabaseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseModelWithDefaults() *DatabaseModel {
	this := DatabaseModel{}
	return &this
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *DatabaseModel) GetOrganization() string {
	if o == nil || IsNil(o.Organization) {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseModel) GetOrganizationOk() (*string, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *DatabaseModel) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *DatabaseModel) SetOrganization(v string) {
	o.Organization = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *DatabaseModel) GetProject() string {
	if o == nil || IsNil(o.Project) {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseModel) GetProjectOk() (*string, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *DatabaseModel) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *DatabaseModel) SetProject(v string) {
	o.Project = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DatabaseModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DatabaseModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DatabaseModel) SetName(v string) {
	o.Name = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *DatabaseModel) GetLabels() map[string]string {
	if o == nil || IsNil(o.Labels) {
		var ret map[string]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseModel) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *DatabaseModel) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *DatabaseModel) SetLabels(v map[string]string) {
	o.Labels = &v
}

// GetTier returns the Tier field value if set, zero value otherwise.
func (o *DatabaseModel) GetTier() string {
	if o == nil || IsNil(o.Tier) {
		var ret string
		return ret
	}
	return *o.Tier
}

// GetTierOk returns a tuple with the Tier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseModel) GetTierOk() (*string, bool) {
	if o == nil || IsNil(o.Tier) {
		return nil, false
	}
	return o.Tier, true
}

// HasTier returns a boolean if a field has been set.
func (o *DatabaseModel) HasTier() bool {
	if o != nil && !IsNil(o.Tier) {
		return true
	}

	return false
}

// SetTier gets a reference to the given string and assigns it to the Tier field.
func (o *DatabaseModel) SetTier(v string) {
	o.Tier = &v
}

// GetMaintenance returns the Maintenance field value if set, zero value otherwise.
func (o *DatabaseModel) GetMaintenance() MaintenanceModel {
	if o == nil || IsNil(o.Maintenance) {
		var ret MaintenanceModel
		return ret
	}
	return *o.Maintenance
}

// GetMaintenanceOk returns a tuple with the Maintenance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseModel) GetMaintenanceOk() (*MaintenanceModel, bool) {
	if o == nil || IsNil(o.Maintenance) {
		return nil, false
	}
	return o.Maintenance, true
}

// HasMaintenance returns a boolean if a field has been set.
func (o *DatabaseModel) HasMaintenance() bool {
	if o != nil && !IsNil(o.Maintenance) {
		return true
	}

	return false
}

// SetMaintenance gets a reference to the given MaintenanceModel and assigns it to the Maintenance field.
func (o *DatabaseModel) SetMaintenance(v MaintenanceModel) {
	o.Maintenance = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *DatabaseModel) GetProperties() DatabasePropertiesModel {
	if o == nil || IsNil(o.Properties) {
		var ret DatabasePropertiesModel
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseModel) GetPropertiesOk() (*DatabasePropertiesModel, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *DatabaseModel) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given DatabasePropertiesModel and assigns it to the Properties field.
func (o *DatabaseModel) SetProperties(v DatabasePropertiesModel) {
	o.Properties = &v
}

// GetResourceVersion returns the ResourceVersion field value if set, zero value otherwise.
func (o *DatabaseModel) GetResourceVersion() string {
	if o == nil || IsNil(o.ResourceVersion) {
		var ret string
		return ret
	}
	return *o.ResourceVersion
}

// GetResourceVersionOk returns a tuple with the ResourceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseModel) GetResourceVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceVersion) {
		return nil, false
	}
	return o.ResourceVersion, true
}

// HasResourceVersion returns a boolean if a field has been set.
func (o *DatabaseModel) HasResourceVersion() bool {
	if o != nil && !IsNil(o.ResourceVersion) {
		return true
	}

	return false
}

// SetResourceVersion gets a reference to the given string and assigns it to the ResourceVersion field.
func (o *DatabaseModel) SetResourceVersion(v string) {
	o.ResourceVersion = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DatabaseModel) GetStatus() DatabaseStatusModel {
	if o == nil || IsNil(o.Status) {
		var ret DatabaseStatusModel
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseModel) GetStatusOk() (*DatabaseStatusModel, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DatabaseModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given DatabaseStatusModel and assigns it to the Status field.
func (o *DatabaseModel) SetStatus(v DatabaseStatusModel) {
	o.Status = &v
}

func (o DatabaseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatabaseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Tier) {
		toSerialize["tier"] = o.Tier
	}
	if !IsNil(o.Maintenance) {
		toSerialize["maintenance"] = o.Maintenance
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.ResourceVersion) {
		toSerialize["resourceVersion"] = o.ResourceVersion
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableDatabaseModel struct {
	value *DatabaseModel
	isSet bool
}

func (v NullableDatabaseModel) Get() *DatabaseModel {
	return v.value
}

func (v *NullableDatabaseModel) Set(val *DatabaseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseModel(val *DatabaseModel) *NullableDatabaseModel {
	return &NullableDatabaseModel{value: val, isSet: true}
}

func (v NullableDatabaseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


