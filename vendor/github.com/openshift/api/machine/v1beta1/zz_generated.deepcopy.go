// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSMachineProviderCondition) DeepCopyInto(out *AWSMachineProviderCondition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSMachineProviderCondition.
func (in *AWSMachineProviderCondition) DeepCopy() *AWSMachineProviderCondition {
	if in == nil {
		return nil
	}
	out := new(AWSMachineProviderCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSMachineProviderConfig) DeepCopyInto(out *AWSMachineProviderConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.AMI.DeepCopyInto(&out.AMI)
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]TagSpecification, len(*in))
		copy(*out, *in)
	}
	if in.IAMInstanceProfile != nil {
		in, out := &in.IAMInstanceProfile, &out.IAMInstanceProfile
		*out = new(AWSResourceReference)
		(*in).DeepCopyInto(*out)
	}
	if in.UserDataSecret != nil {
		in, out := &in.UserDataSecret, &out.UserDataSecret
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.CredentialsSecret != nil {
		in, out := &in.CredentialsSecret, &out.CredentialsSecret
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.KeyName != nil {
		in, out := &in.KeyName, &out.KeyName
		*out = new(string)
		**out = **in
	}
	if in.PublicIP != nil {
		in, out := &in.PublicIP, &out.PublicIP
		*out = new(bool)
		**out = **in
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]AWSResourceReference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Subnet.DeepCopyInto(&out.Subnet)
	out.Placement = in.Placement
	if in.LoadBalancers != nil {
		in, out := &in.LoadBalancers, &out.LoadBalancers
		*out = make([]LoadBalancerReference, len(*in))
		copy(*out, *in)
	}
	if in.BlockDevices != nil {
		in, out := &in.BlockDevices, &out.BlockDevices
		*out = make([]BlockDeviceMappingSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SpotMarketOptions != nil {
		in, out := &in.SpotMarketOptions, &out.SpotMarketOptions
		*out = new(SpotMarketOptions)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSMachineProviderConfig.
func (in *AWSMachineProviderConfig) DeepCopy() *AWSMachineProviderConfig {
	if in == nil {
		return nil
	}
	out := new(AWSMachineProviderConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSMachineProviderConfigList) DeepCopyInto(out *AWSMachineProviderConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AWSMachineProviderConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSMachineProviderConfigList.
func (in *AWSMachineProviderConfigList) DeepCopy() *AWSMachineProviderConfigList {
	if in == nil {
		return nil
	}
	out := new(AWSMachineProviderConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSMachineProviderStatus) DeepCopyInto(out *AWSMachineProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.InstanceID != nil {
		in, out := &in.InstanceID, &out.InstanceID
		*out = new(string)
		**out = **in
	}
	if in.InstanceState != nil {
		in, out := &in.InstanceState, &out.InstanceState
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]AWSMachineProviderCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSMachineProviderStatus.
func (in *AWSMachineProviderStatus) DeepCopy() *AWSMachineProviderStatus {
	if in == nil {
		return nil
	}
	out := new(AWSMachineProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSResourceReference) DeepCopyInto(out *AWSResourceReference) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(string)
		**out = **in
	}
	if in.Filters != nil {
		in, out := &in.Filters, &out.Filters
		*out = make([]Filter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSResourceReference.
func (in *AWSResourceReference) DeepCopy() *AWSResourceReference {
	if in == nil {
		return nil
	}
	out := new(AWSResourceReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureMachineProviderCondition) DeepCopyInto(out *AzureMachineProviderCondition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureMachineProviderCondition.
func (in *AzureMachineProviderCondition) DeepCopy() *AzureMachineProviderCondition {
	if in == nil {
		return nil
	}
	out := new(AzureMachineProviderCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureMachineProviderSpec) DeepCopyInto(out *AzureMachineProviderSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.UserDataSecret != nil {
		in, out := &in.UserDataSecret, &out.UserDataSecret
		*out = new(v1.SecretReference)
		**out = **in
	}
	if in.CredentialsSecret != nil {
		in, out := &in.CredentialsSecret, &out.CredentialsSecret
		*out = new(v1.SecretReference)
		**out = **in
	}
	out.Image = in.Image
	in.OSDisk.DeepCopyInto(&out.OSDisk)
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ApplicationSecurityGroups != nil {
		in, out := &in.ApplicationSecurityGroups, &out.ApplicationSecurityGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NatRule != nil {
		in, out := &in.NatRule, &out.NatRule
		*out = new(int64)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
	if in.SpotVMOptions != nil {
		in, out := &in.SpotVMOptions, &out.SpotVMOptions
		*out = new(SpotVMOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.SecurityProfile != nil {
		in, out := &in.SecurityProfile, &out.SecurityProfile
		*out = new(SecurityProfile)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureMachineProviderSpec.
func (in *AzureMachineProviderSpec) DeepCopy() *AzureMachineProviderSpec {
	if in == nil {
		return nil
	}
	out := new(AzureMachineProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureMachineProviderStatus) DeepCopyInto(out *AzureMachineProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.VMID != nil {
		in, out := &in.VMID, &out.VMID
		*out = new(string)
		**out = **in
	}
	if in.VMState != nil {
		in, out := &in.VMState, &out.VMState
		*out = new(AzureVMState)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]AzureMachineProviderCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureMachineProviderStatus.
func (in *AzureMachineProviderStatus) DeepCopy() *AzureMachineProviderStatus {
	if in == nil {
		return nil
	}
	out := new(AzureMachineProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockDeviceMappingSpec) DeepCopyInto(out *BlockDeviceMappingSpec) {
	*out = *in
	if in.DeviceName != nil {
		in, out := &in.DeviceName, &out.DeviceName
		*out = new(string)
		**out = **in
	}
	if in.EBS != nil {
		in, out := &in.EBS, &out.EBS
		*out = new(EBSBlockDeviceSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.NoDevice != nil {
		in, out := &in.NoDevice, &out.NoDevice
		*out = new(string)
		**out = **in
	}
	if in.VirtualName != nil {
		in, out := &in.VirtualName, &out.VirtualName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockDeviceMappingSpec.
func (in *BlockDeviceMappingSpec) DeepCopy() *BlockDeviceMappingSpec {
	if in == nil {
		return nil
	}
	out := new(BlockDeviceMappingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Conditions) DeepCopyInto(out *Conditions) {
	{
		in := &in
		*out = make(Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Conditions.
func (in Conditions) DeepCopy() Conditions {
	if in == nil {
		return nil
	}
	out := new(Conditions)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskEncryptionSetParameters) DeepCopyInto(out *DiskEncryptionSetParameters) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskEncryptionSetParameters.
func (in *DiskEncryptionSetParameters) DeepCopy() *DiskEncryptionSetParameters {
	if in == nil {
		return nil
	}
	out := new(DiskEncryptionSetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EBSBlockDeviceSpec) DeepCopyInto(out *EBSBlockDeviceSpec) {
	*out = *in
	if in.DeleteOnTermination != nil {
		in, out := &in.DeleteOnTermination, &out.DeleteOnTermination
		*out = new(bool)
		**out = **in
	}
	if in.Encrypted != nil {
		in, out := &in.Encrypted, &out.Encrypted
		*out = new(bool)
		**out = **in
	}
	in.KMSKey.DeepCopyInto(&out.KMSKey)
	if in.Iops != nil {
		in, out := &in.Iops, &out.Iops
		*out = new(int64)
		**out = **in
	}
	if in.VolumeSize != nil {
		in, out := &in.VolumeSize, &out.VolumeSize
		*out = new(int64)
		**out = **in
	}
	if in.VolumeType != nil {
		in, out := &in.VolumeType, &out.VolumeType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EBSBlockDeviceSpec.
func (in *EBSBlockDeviceSpec) DeepCopy() *EBSBlockDeviceSpec {
	if in == nil {
		return nil
	}
	out := new(EBSBlockDeviceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Filter) DeepCopyInto(out *Filter) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Filter.
func (in *Filter) DeepCopy() *Filter {
	if in == nil {
		return nil
	}
	out := new(Filter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPDisk) DeepCopyInto(out *GCPDisk) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.EncryptionKey != nil {
		in, out := &in.EncryptionKey, &out.EncryptionKey
		*out = new(GCPEncryptionKeyReference)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPDisk.
func (in *GCPDisk) DeepCopy() *GCPDisk {
	if in == nil {
		return nil
	}
	out := new(GCPDisk)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPEncryptionKeyReference) DeepCopyInto(out *GCPEncryptionKeyReference) {
	*out = *in
	if in.KMSKey != nil {
		in, out := &in.KMSKey, &out.KMSKey
		*out = new(GCPKMSKeyReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPEncryptionKeyReference.
func (in *GCPEncryptionKeyReference) DeepCopy() *GCPEncryptionKeyReference {
	if in == nil {
		return nil
	}
	out := new(GCPEncryptionKeyReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPKMSKeyReference) DeepCopyInto(out *GCPKMSKeyReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPKMSKeyReference.
func (in *GCPKMSKeyReference) DeepCopy() *GCPKMSKeyReference {
	if in == nil {
		return nil
	}
	out := new(GCPKMSKeyReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPMachineProviderCondition) DeepCopyInto(out *GCPMachineProviderCondition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPMachineProviderCondition.
func (in *GCPMachineProviderCondition) DeepCopy() *GCPMachineProviderCondition {
	if in == nil {
		return nil
	}
	out := new(GCPMachineProviderCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPMachineProviderSpec) DeepCopyInto(out *GCPMachineProviderSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.UserDataSecret != nil {
		in, out := &in.UserDataSecret, &out.UserDataSecret
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.CredentialsSecret != nil {
		in, out := &in.CredentialsSecret, &out.CredentialsSecret
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.Disks != nil {
		in, out := &in.Disks, &out.Disks
		*out = make([]*GCPDisk, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(GCPDisk)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make([]*GCPMetadata, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(GCPMetadata)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.NetworkInterfaces != nil {
		in, out := &in.NetworkInterfaces, &out.NetworkInterfaces
		*out = make([]*GCPNetworkInterface, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(GCPNetworkInterface)
				**out = **in
			}
		}
	}
	if in.ServiceAccounts != nil {
		in, out := &in.ServiceAccounts, &out.ServiceAccounts
		*out = make([]GCPServiceAccount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TargetPools != nil {
		in, out := &in.TargetPools, &out.TargetPools
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPMachineProviderSpec.
func (in *GCPMachineProviderSpec) DeepCopy() *GCPMachineProviderSpec {
	if in == nil {
		return nil
	}
	out := new(GCPMachineProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPMachineProviderStatus) DeepCopyInto(out *GCPMachineProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.InstanceID != nil {
		in, out := &in.InstanceID, &out.InstanceID
		*out = new(string)
		**out = **in
	}
	if in.InstanceState != nil {
		in, out := &in.InstanceState, &out.InstanceState
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]GCPMachineProviderCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPMachineProviderStatus.
func (in *GCPMachineProviderStatus) DeepCopy() *GCPMachineProviderStatus {
	if in == nil {
		return nil
	}
	out := new(GCPMachineProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPMetadata) DeepCopyInto(out *GCPMetadata) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPMetadata.
func (in *GCPMetadata) DeepCopy() *GCPMetadata {
	if in == nil {
		return nil
	}
	out := new(GCPMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPNetworkInterface) DeepCopyInto(out *GCPNetworkInterface) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPNetworkInterface.
func (in *GCPNetworkInterface) DeepCopy() *GCPNetworkInterface {
	if in == nil {
		return nil
	}
	out := new(GCPNetworkInterface)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPServiceAccount) DeepCopyInto(out *GCPServiceAccount) {
	*out = *in
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPServiceAccount.
func (in *GCPServiceAccount) DeepCopy() *GCPServiceAccount {
	if in == nil {
		return nil
	}
	out := new(GCPServiceAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image) DeepCopyInto(out *Image) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image.
func (in *Image) DeepCopy() *Image {
	if in == nil {
		return nil
	}
	out := new(Image)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LastOperation) DeepCopyInto(out *LastOperation) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.LastUpdated != nil {
		in, out := &in.LastUpdated, &out.LastUpdated
		*out = (*in).DeepCopy()
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LastOperation.
func (in *LastOperation) DeepCopy() *LastOperation {
	if in == nil {
		return nil
	}
	out := new(LastOperation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoadBalancerReference) DeepCopyInto(out *LoadBalancerReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoadBalancerReference.
func (in *LoadBalancerReference) DeepCopy() *LoadBalancerReference {
	if in == nil {
		return nil
	}
	out := new(LoadBalancerReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Machine) DeepCopyInto(out *Machine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Machine.
func (in *Machine) DeepCopy() *Machine {
	if in == nil {
		return nil
	}
	out := new(Machine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Machine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineList) DeepCopyInto(out *MachineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Machine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineList.
func (in *MachineList) DeepCopy() *MachineList {
	if in == nil {
		return nil
	}
	out := new(MachineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MachineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineSpec) DeepCopyInto(out *MachineSpec) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Taints != nil {
		in, out := &in.Taints, &out.Taints
		*out = make([]v1.Taint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.ProviderSpec.DeepCopyInto(&out.ProviderSpec)
	if in.ProviderID != nil {
		in, out := &in.ProviderID, &out.ProviderID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineSpec.
func (in *MachineSpec) DeepCopy() *MachineSpec {
	if in == nil {
		return nil
	}
	out := new(MachineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineStatus) DeepCopyInto(out *MachineStatus) {
	*out = *in
	if in.NodeRef != nil {
		in, out := &in.NodeRef, &out.NodeRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.LastUpdated != nil {
		in, out := &in.LastUpdated, &out.LastUpdated
		*out = (*in).DeepCopy()
	}
	if in.ErrorReason != nil {
		in, out := &in.ErrorReason, &out.ErrorReason
		*out = new(MachineStatusError)
		**out = **in
	}
	if in.ErrorMessage != nil {
		in, out := &in.ErrorMessage, &out.ErrorMessage
		*out = new(string)
		**out = **in
	}
	if in.ProviderStatus != nil {
		in, out := &in.ProviderStatus, &out.ProviderStatus
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]v1.NodeAddress, len(*in))
		copy(*out, *in)
	}
	if in.LastOperation != nil {
		in, out := &in.LastOperation, &out.LastOperation
		*out = new(LastOperation)
		(*in).DeepCopyInto(*out)
	}
	if in.Phase != nil {
		in, out := &in.Phase, &out.Phase
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineStatus.
func (in *MachineStatus) DeepCopy() *MachineStatus {
	if in == nil {
		return nil
	}
	out := new(MachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedDiskParameters) DeepCopyInto(out *ManagedDiskParameters) {
	*out = *in
	if in.DiskEncryptionSet != nil {
		in, out := &in.DiskEncryptionSet, &out.DiskEncryptionSet
		*out = new(DiskEncryptionSetParameters)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedDiskParameters.
func (in *ManagedDiskParameters) DeepCopy() *ManagedDiskParameters {
	if in == nil {
		return nil
	}
	out := new(ManagedDiskParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkDeviceSpec) DeepCopyInto(out *NetworkDeviceSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkDeviceSpec.
func (in *NetworkDeviceSpec) DeepCopy() *NetworkDeviceSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkDeviceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkSpec) DeepCopyInto(out *NetworkSpec) {
	*out = *in
	if in.Devices != nil {
		in, out := &in.Devices, &out.Devices
		*out = make([]NetworkDeviceSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkSpec.
func (in *NetworkSpec) DeepCopy() *NetworkSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OSDisk) DeepCopyInto(out *OSDisk) {
	*out = *in
	in.ManagedDisk.DeepCopyInto(&out.ManagedDisk)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OSDisk.
func (in *OSDisk) DeepCopy() *OSDisk {
	if in == nil {
		return nil
	}
	out := new(OSDisk)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectMeta) DeepCopyInto(out *ObjectMeta) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.OwnerReferences != nil {
		in, out := &in.OwnerReferences, &out.OwnerReferences
		*out = make([]metav1.OwnerReference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectMeta.
func (in *ObjectMeta) DeepCopy() *ObjectMeta {
	if in == nil {
		return nil
	}
	out := new(ObjectMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Placement) DeepCopyInto(out *Placement) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Placement.
func (in *Placement) DeepCopy() *Placement {
	if in == nil {
		return nil
	}
	out := new(Placement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderSpec) DeepCopyInto(out *ProviderSpec) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderSpec.
func (in *ProviderSpec) DeepCopy() *ProviderSpec {
	if in == nil {
		return nil
	}
	out := new(ProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityProfile) DeepCopyInto(out *SecurityProfile) {
	*out = *in
	if in.EncryptionAtHost != nil {
		in, out := &in.EncryptionAtHost, &out.EncryptionAtHost
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityProfile.
func (in *SecurityProfile) DeepCopy() *SecurityProfile {
	if in == nil {
		return nil
	}
	out := new(SecurityProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpotMarketOptions) DeepCopyInto(out *SpotMarketOptions) {
	*out = *in
	if in.MaxPrice != nil {
		in, out := &in.MaxPrice, &out.MaxPrice
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpotMarketOptions.
func (in *SpotMarketOptions) DeepCopy() *SpotMarketOptions {
	if in == nil {
		return nil
	}
	out := new(SpotMarketOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpotVMOptions) DeepCopyInto(out *SpotVMOptions) {
	*out = *in
	if in.MaxPrice != nil {
		in, out := &in.MaxPrice, &out.MaxPrice
		x := (*in).DeepCopy()
		*out = &x
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpotVMOptions.
func (in *SpotVMOptions) DeepCopy() *SpotVMOptions {
	if in == nil {
		return nil
	}
	out := new(SpotVMOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagSpecification) DeepCopyInto(out *TagSpecification) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagSpecification.
func (in *TagSpecification) DeepCopy() *TagSpecification {
	if in == nil {
		return nil
	}
	out := new(TagSpecification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSphereMachineProviderCondition) DeepCopyInto(out *VSphereMachineProviderCondition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSphereMachineProviderCondition.
func (in *VSphereMachineProviderCondition) DeepCopy() *VSphereMachineProviderCondition {
	if in == nil {
		return nil
	}
	out := new(VSphereMachineProviderCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSphereMachineProviderSpec) DeepCopyInto(out *VSphereMachineProviderSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.UserDataSecret != nil {
		in, out := &in.UserDataSecret, &out.UserDataSecret
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.CredentialsSecret != nil {
		in, out := &in.CredentialsSecret, &out.CredentialsSecret
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.Workspace != nil {
		in, out := &in.Workspace, &out.Workspace
		*out = new(Workspace)
		**out = **in
	}
	in.Network.DeepCopyInto(&out.Network)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSphereMachineProviderSpec.
func (in *VSphereMachineProviderSpec) DeepCopy() *VSphereMachineProviderSpec {
	if in == nil {
		return nil
	}
	out := new(VSphereMachineProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VSphereMachineProviderStatus) DeepCopyInto(out *VSphereMachineProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.InstanceID != nil {
		in, out := &in.InstanceID, &out.InstanceID
		*out = new(string)
		**out = **in
	}
	if in.InstanceState != nil {
		in, out := &in.InstanceState, &out.InstanceState
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]VSphereMachineProviderCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VSphereMachineProviderStatus.
func (in *VSphereMachineProviderStatus) DeepCopy() *VSphereMachineProviderStatus {
	if in == nil {
		return nil
	}
	out := new(VSphereMachineProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Workspace) DeepCopyInto(out *Workspace) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Workspace.
func (in *Workspace) DeepCopy() *Workspace {
	if in == nil {
		return nil
	}
	out := new(Workspace)
	in.DeepCopyInto(out)
	return out
}
