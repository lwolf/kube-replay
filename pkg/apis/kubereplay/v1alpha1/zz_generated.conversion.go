// +build !ignore_autogenerated

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1alpha1

import (
	"github.com/lwolf/kubereplay/pkg/apis/kubereplay"
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/runtime"
	"unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_ElasticsearchSilo_To_kubereplay_ElasticsearchSilo,
		Convert_kubereplay_ElasticsearchSilo_To_v1alpha1_ElasticsearchSilo,
		Convert_v1alpha1_FileSilo_To_kubereplay_FileSilo,
		Convert_kubereplay_FileSilo_To_v1alpha1_FileSilo,
		Convert_v1alpha1_Harvester_To_kubereplay_Harvester,
		Convert_kubereplay_Harvester_To_v1alpha1_Harvester,
		Convert_v1alpha1_HarvesterList_To_kubereplay_HarvesterList,
		Convert_kubereplay_HarvesterList_To_v1alpha1_HarvesterList,
		Convert_v1alpha1_HarvesterSpec_To_kubereplay_HarvesterSpec,
		Convert_kubereplay_HarvesterSpec_To_v1alpha1_HarvesterSpec,
		Convert_v1alpha1_HarvesterStatus_To_kubereplay_HarvesterStatus,
		Convert_kubereplay_HarvesterStatus_To_v1alpha1_HarvesterStatus,
		Convert_v1alpha1_HarvesterStatusStrategy_To_kubereplay_HarvesterStatusStrategy,
		Convert_kubereplay_HarvesterStatusStrategy_To_v1alpha1_HarvesterStatusStrategy,
		Convert_v1alpha1_HarvesterStrategy_To_kubereplay_HarvesterStrategy,
		Convert_kubereplay_HarvesterStrategy_To_v1alpha1_HarvesterStrategy,
		Convert_v1alpha1_HttpSilo_To_kubereplay_HttpSilo,
		Convert_kubereplay_HttpSilo_To_v1alpha1_HttpSilo,
		Convert_v1alpha1_KafkaSilo_To_kubereplay_KafkaSilo,
		Convert_kubereplay_KafkaSilo_To_v1alpha1_KafkaSilo,
		Convert_v1alpha1_Refinery_To_kubereplay_Refinery,
		Convert_kubereplay_Refinery_To_v1alpha1_Refinery,
		Convert_v1alpha1_RefineryList_To_kubereplay_RefineryList,
		Convert_kubereplay_RefineryList_To_v1alpha1_RefineryList,
		Convert_v1alpha1_RefinerySpec_To_kubereplay_RefinerySpec,
		Convert_kubereplay_RefinerySpec_To_v1alpha1_RefinerySpec,
		Convert_v1alpha1_RefineryStatus_To_kubereplay_RefineryStatus,
		Convert_kubereplay_RefineryStatus_To_v1alpha1_RefineryStatus,
		Convert_v1alpha1_RefineryStatusStrategy_To_kubereplay_RefineryStatusStrategy,
		Convert_kubereplay_RefineryStatusStrategy_To_v1alpha1_RefineryStatusStrategy,
		Convert_v1alpha1_RefineryStorage_To_kubereplay_RefineryStorage,
		Convert_kubereplay_RefineryStorage_To_v1alpha1_RefineryStorage,
		Convert_v1alpha1_RefineryStrategy_To_kubereplay_RefineryStrategy,
		Convert_kubereplay_RefineryStrategy_To_v1alpha1_RefineryStrategy,
		Convert_v1alpha1_StdoutSilo_To_kubereplay_StdoutSilo,
		Convert_kubereplay_StdoutSilo_To_v1alpha1_StdoutSilo,
		Convert_v1alpha1_TcpSilo_To_kubereplay_TcpSilo,
		Convert_kubereplay_TcpSilo_To_v1alpha1_TcpSilo,
	)
}

func autoConvert_v1alpha1_ElasticsearchSilo_To_kubereplay_ElasticsearchSilo(in *ElasticsearchSilo, out *kubereplay.ElasticsearchSilo, s conversion.Scope) error {
	out.Enabled = in.Enabled
	out.Uri = in.Uri
	return nil
}

// Convert_v1alpha1_ElasticsearchSilo_To_kubereplay_ElasticsearchSilo is an autogenerated conversion function.
func Convert_v1alpha1_ElasticsearchSilo_To_kubereplay_ElasticsearchSilo(in *ElasticsearchSilo, out *kubereplay.ElasticsearchSilo, s conversion.Scope) error {
	return autoConvert_v1alpha1_ElasticsearchSilo_To_kubereplay_ElasticsearchSilo(in, out, s)
}

func autoConvert_kubereplay_ElasticsearchSilo_To_v1alpha1_ElasticsearchSilo(in *kubereplay.ElasticsearchSilo, out *ElasticsearchSilo, s conversion.Scope) error {
	out.Enabled = in.Enabled
	out.Uri = in.Uri
	return nil
}

// Convert_kubereplay_ElasticsearchSilo_To_v1alpha1_ElasticsearchSilo is an autogenerated conversion function.
func Convert_kubereplay_ElasticsearchSilo_To_v1alpha1_ElasticsearchSilo(in *kubereplay.ElasticsearchSilo, out *ElasticsearchSilo, s conversion.Scope) error {
	return autoConvert_kubereplay_ElasticsearchSilo_To_v1alpha1_ElasticsearchSilo(in, out, s)
}

func autoConvert_v1alpha1_FileSilo_To_kubereplay_FileSilo(in *FileSilo, out *kubereplay.FileSilo, s conversion.Scope) error {
	out.Enabled = in.Enabled
	out.Filename = in.Filename
	out.Append = in.Append
	out.FlushInterval = in.FlushInterval
	out.QueueSize = in.QueueSize
	out.FileLimit = in.FileLimit
	return nil
}

// Convert_v1alpha1_FileSilo_To_kubereplay_FileSilo is an autogenerated conversion function.
func Convert_v1alpha1_FileSilo_To_kubereplay_FileSilo(in *FileSilo, out *kubereplay.FileSilo, s conversion.Scope) error {
	return autoConvert_v1alpha1_FileSilo_To_kubereplay_FileSilo(in, out, s)
}

func autoConvert_kubereplay_FileSilo_To_v1alpha1_FileSilo(in *kubereplay.FileSilo, out *FileSilo, s conversion.Scope) error {
	out.Enabled = in.Enabled
	out.Filename = in.Filename
	out.Append = in.Append
	out.FlushInterval = in.FlushInterval
	out.QueueSize = in.QueueSize
	out.FileLimit = in.FileLimit
	return nil
}

// Convert_kubereplay_FileSilo_To_v1alpha1_FileSilo is an autogenerated conversion function.
func Convert_kubereplay_FileSilo_To_v1alpha1_FileSilo(in *kubereplay.FileSilo, out *FileSilo, s conversion.Scope) error {
	return autoConvert_kubereplay_FileSilo_To_v1alpha1_FileSilo(in, out, s)
}

func autoConvert_v1alpha1_Harvester_To_kubereplay_Harvester(in *Harvester, out *kubereplay.Harvester, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_HarvesterSpec_To_kubereplay_HarvesterSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_HarvesterStatus_To_kubereplay_HarvesterStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_Harvester_To_kubereplay_Harvester is an autogenerated conversion function.
func Convert_v1alpha1_Harvester_To_kubereplay_Harvester(in *Harvester, out *kubereplay.Harvester, s conversion.Scope) error {
	return autoConvert_v1alpha1_Harvester_To_kubereplay_Harvester(in, out, s)
}

func autoConvert_kubereplay_Harvester_To_v1alpha1_Harvester(in *kubereplay.Harvester, out *Harvester, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_kubereplay_HarvesterSpec_To_v1alpha1_HarvesterSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_kubereplay_HarvesterStatus_To_v1alpha1_HarvesterStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_kubereplay_Harvester_To_v1alpha1_Harvester is an autogenerated conversion function.
func Convert_kubereplay_Harvester_To_v1alpha1_Harvester(in *kubereplay.Harvester, out *Harvester, s conversion.Scope) error {
	return autoConvert_kubereplay_Harvester_To_v1alpha1_Harvester(in, out, s)
}

func autoConvert_v1alpha1_HarvesterList_To_kubereplay_HarvesterList(in *HarvesterList, out *kubereplay.HarvesterList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]kubereplay.Harvester)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_HarvesterList_To_kubereplay_HarvesterList is an autogenerated conversion function.
func Convert_v1alpha1_HarvesterList_To_kubereplay_HarvesterList(in *HarvesterList, out *kubereplay.HarvesterList, s conversion.Scope) error {
	return autoConvert_v1alpha1_HarvesterList_To_kubereplay_HarvesterList(in, out, s)
}

func autoConvert_kubereplay_HarvesterList_To_v1alpha1_HarvesterList(in *kubereplay.HarvesterList, out *HarvesterList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Harvester)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_kubereplay_HarvesterList_To_v1alpha1_HarvesterList is an autogenerated conversion function.
func Convert_kubereplay_HarvesterList_To_v1alpha1_HarvesterList(in *kubereplay.HarvesterList, out *HarvesterList, s conversion.Scope) error {
	return autoConvert_kubereplay_HarvesterList_To_v1alpha1_HarvesterList(in, out, s)
}

func autoConvert_v1alpha1_HarvesterSpec_To_kubereplay_HarvesterSpec(in *HarvesterSpec, out *kubereplay.HarvesterSpec, s conversion.Scope) error {
	out.Selector = *(*map[string]string)(unsafe.Pointer(&in.Selector))
	out.AppPort = in.AppPort
	out.Refinery = in.Refinery
	out.SegmentSize = in.SegmentSize
	return nil
}

// Convert_v1alpha1_HarvesterSpec_To_kubereplay_HarvesterSpec is an autogenerated conversion function.
func Convert_v1alpha1_HarvesterSpec_To_kubereplay_HarvesterSpec(in *HarvesterSpec, out *kubereplay.HarvesterSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_HarvesterSpec_To_kubereplay_HarvesterSpec(in, out, s)
}

func autoConvert_kubereplay_HarvesterSpec_To_v1alpha1_HarvesterSpec(in *kubereplay.HarvesterSpec, out *HarvesterSpec, s conversion.Scope) error {
	out.Selector = *(*map[string]string)(unsafe.Pointer(&in.Selector))
	out.AppPort = in.AppPort
	out.Refinery = in.Refinery
	out.SegmentSize = in.SegmentSize
	return nil
}

// Convert_kubereplay_HarvesterSpec_To_v1alpha1_HarvesterSpec is an autogenerated conversion function.
func Convert_kubereplay_HarvesterSpec_To_v1alpha1_HarvesterSpec(in *kubereplay.HarvesterSpec, out *HarvesterSpec, s conversion.Scope) error {
	return autoConvert_kubereplay_HarvesterSpec_To_v1alpha1_HarvesterSpec(in, out, s)
}

func autoConvert_v1alpha1_HarvesterStatus_To_kubereplay_HarvesterStatus(in *HarvesterStatus, out *kubereplay.HarvesterStatus, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha1_HarvesterStatus_To_kubereplay_HarvesterStatus is an autogenerated conversion function.
func Convert_v1alpha1_HarvesterStatus_To_kubereplay_HarvesterStatus(in *HarvesterStatus, out *kubereplay.HarvesterStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_HarvesterStatus_To_kubereplay_HarvesterStatus(in, out, s)
}

func autoConvert_kubereplay_HarvesterStatus_To_v1alpha1_HarvesterStatus(in *kubereplay.HarvesterStatus, out *HarvesterStatus, s conversion.Scope) error {
	return nil
}

// Convert_kubereplay_HarvesterStatus_To_v1alpha1_HarvesterStatus is an autogenerated conversion function.
func Convert_kubereplay_HarvesterStatus_To_v1alpha1_HarvesterStatus(in *kubereplay.HarvesterStatus, out *HarvesterStatus, s conversion.Scope) error {
	return autoConvert_kubereplay_HarvesterStatus_To_v1alpha1_HarvesterStatus(in, out, s)
}

func autoConvert_v1alpha1_HarvesterStatusStrategy_To_kubereplay_HarvesterStatusStrategy(in *HarvesterStatusStrategy, out *kubereplay.HarvesterStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_v1alpha1_HarvesterStatusStrategy_To_kubereplay_HarvesterStatusStrategy is an autogenerated conversion function.
func Convert_v1alpha1_HarvesterStatusStrategy_To_kubereplay_HarvesterStatusStrategy(in *HarvesterStatusStrategy, out *kubereplay.HarvesterStatusStrategy, s conversion.Scope) error {
	return autoConvert_v1alpha1_HarvesterStatusStrategy_To_kubereplay_HarvesterStatusStrategy(in, out, s)
}

func autoConvert_kubereplay_HarvesterStatusStrategy_To_v1alpha1_HarvesterStatusStrategy(in *kubereplay.HarvesterStatusStrategy, out *HarvesterStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_kubereplay_HarvesterStatusStrategy_To_v1alpha1_HarvesterStatusStrategy is an autogenerated conversion function.
func Convert_kubereplay_HarvesterStatusStrategy_To_v1alpha1_HarvesterStatusStrategy(in *kubereplay.HarvesterStatusStrategy, out *HarvesterStatusStrategy, s conversion.Scope) error {
	return autoConvert_kubereplay_HarvesterStatusStrategy_To_v1alpha1_HarvesterStatusStrategy(in, out, s)
}

func autoConvert_v1alpha1_HarvesterStrategy_To_kubereplay_HarvesterStrategy(in *HarvesterStrategy, out *kubereplay.HarvesterStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_v1alpha1_HarvesterStrategy_To_kubereplay_HarvesterStrategy is an autogenerated conversion function.
func Convert_v1alpha1_HarvesterStrategy_To_kubereplay_HarvesterStrategy(in *HarvesterStrategy, out *kubereplay.HarvesterStrategy, s conversion.Scope) error {
	return autoConvert_v1alpha1_HarvesterStrategy_To_kubereplay_HarvesterStrategy(in, out, s)
}

func autoConvert_kubereplay_HarvesterStrategy_To_v1alpha1_HarvesterStrategy(in *kubereplay.HarvesterStrategy, out *HarvesterStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_kubereplay_HarvesterStrategy_To_v1alpha1_HarvesterStrategy is an autogenerated conversion function.
func Convert_kubereplay_HarvesterStrategy_To_v1alpha1_HarvesterStrategy(in *kubereplay.HarvesterStrategy, out *HarvesterStrategy, s conversion.Scope) error {
	return autoConvert_kubereplay_HarvesterStrategy_To_v1alpha1_HarvesterStrategy(in, out, s)
}

func autoConvert_v1alpha1_HttpSilo_To_kubereplay_HttpSilo(in *HttpSilo, out *kubereplay.HttpSilo, s conversion.Scope) error {
	out.Enabled = in.Enabled
	out.Uri = in.Uri
	out.Debug = in.Debug
	out.ResponseBuffer = in.ResponseBuffer
	return nil
}

// Convert_v1alpha1_HttpSilo_To_kubereplay_HttpSilo is an autogenerated conversion function.
func Convert_v1alpha1_HttpSilo_To_kubereplay_HttpSilo(in *HttpSilo, out *kubereplay.HttpSilo, s conversion.Scope) error {
	return autoConvert_v1alpha1_HttpSilo_To_kubereplay_HttpSilo(in, out, s)
}

func autoConvert_kubereplay_HttpSilo_To_v1alpha1_HttpSilo(in *kubereplay.HttpSilo, out *HttpSilo, s conversion.Scope) error {
	out.Enabled = in.Enabled
	out.Uri = in.Uri
	out.Debug = in.Debug
	out.ResponseBuffer = in.ResponseBuffer
	return nil
}

// Convert_kubereplay_HttpSilo_To_v1alpha1_HttpSilo is an autogenerated conversion function.
func Convert_kubereplay_HttpSilo_To_v1alpha1_HttpSilo(in *kubereplay.HttpSilo, out *HttpSilo, s conversion.Scope) error {
	return autoConvert_kubereplay_HttpSilo_To_v1alpha1_HttpSilo(in, out, s)
}

func autoConvert_v1alpha1_KafkaSilo_To_kubereplay_KafkaSilo(in *KafkaSilo, out *kubereplay.KafkaSilo, s conversion.Scope) error {
	out.Enabled = in.Enabled
	out.Uri = in.Uri
	out.Json = in.Json
	out.Topic = in.Topic
	return nil
}

// Convert_v1alpha1_KafkaSilo_To_kubereplay_KafkaSilo is an autogenerated conversion function.
func Convert_v1alpha1_KafkaSilo_To_kubereplay_KafkaSilo(in *KafkaSilo, out *kubereplay.KafkaSilo, s conversion.Scope) error {
	return autoConvert_v1alpha1_KafkaSilo_To_kubereplay_KafkaSilo(in, out, s)
}

func autoConvert_kubereplay_KafkaSilo_To_v1alpha1_KafkaSilo(in *kubereplay.KafkaSilo, out *KafkaSilo, s conversion.Scope) error {
	out.Enabled = in.Enabled
	out.Uri = in.Uri
	out.Json = in.Json
	out.Topic = in.Topic
	return nil
}

// Convert_kubereplay_KafkaSilo_To_v1alpha1_KafkaSilo is an autogenerated conversion function.
func Convert_kubereplay_KafkaSilo_To_v1alpha1_KafkaSilo(in *kubereplay.KafkaSilo, out *KafkaSilo, s conversion.Scope) error {
	return autoConvert_kubereplay_KafkaSilo_To_v1alpha1_KafkaSilo(in, out, s)
}

func autoConvert_v1alpha1_Refinery_To_kubereplay_Refinery(in *Refinery, out *kubereplay.Refinery, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_RefinerySpec_To_kubereplay_RefinerySpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_RefineryStatus_To_kubereplay_RefineryStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_Refinery_To_kubereplay_Refinery is an autogenerated conversion function.
func Convert_v1alpha1_Refinery_To_kubereplay_Refinery(in *Refinery, out *kubereplay.Refinery, s conversion.Scope) error {
	return autoConvert_v1alpha1_Refinery_To_kubereplay_Refinery(in, out, s)
}

func autoConvert_kubereplay_Refinery_To_v1alpha1_Refinery(in *kubereplay.Refinery, out *Refinery, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_kubereplay_RefinerySpec_To_v1alpha1_RefinerySpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_kubereplay_RefineryStatus_To_v1alpha1_RefineryStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_kubereplay_Refinery_To_v1alpha1_Refinery is an autogenerated conversion function.
func Convert_kubereplay_Refinery_To_v1alpha1_Refinery(in *kubereplay.Refinery, out *Refinery, s conversion.Scope) error {
	return autoConvert_kubereplay_Refinery_To_v1alpha1_Refinery(in, out, s)
}

func autoConvert_v1alpha1_RefineryList_To_kubereplay_RefineryList(in *RefineryList, out *kubereplay.RefineryList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]kubereplay.Refinery)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_RefineryList_To_kubereplay_RefineryList is an autogenerated conversion function.
func Convert_v1alpha1_RefineryList_To_kubereplay_RefineryList(in *RefineryList, out *kubereplay.RefineryList, s conversion.Scope) error {
	return autoConvert_v1alpha1_RefineryList_To_kubereplay_RefineryList(in, out, s)
}

func autoConvert_kubereplay_RefineryList_To_v1alpha1_RefineryList(in *kubereplay.RefineryList, out *RefineryList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Refinery)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_kubereplay_RefineryList_To_v1alpha1_RefineryList is an autogenerated conversion function.
func Convert_kubereplay_RefineryList_To_v1alpha1_RefineryList(in *kubereplay.RefineryList, out *RefineryList, s conversion.Scope) error {
	return autoConvert_kubereplay_RefineryList_To_v1alpha1_RefineryList(in, out, s)
}

func autoConvert_v1alpha1_RefinerySpec_To_kubereplay_RefinerySpec(in *RefinerySpec, out *kubereplay.RefinerySpec, s conversion.Scope) error {
	out.Workers = in.Workers
	out.Timeout = in.Timeout
	out.Storage = (*kubereplay.RefineryStorage)(unsafe.Pointer(in.Storage))
	return nil
}

// Convert_v1alpha1_RefinerySpec_To_kubereplay_RefinerySpec is an autogenerated conversion function.
func Convert_v1alpha1_RefinerySpec_To_kubereplay_RefinerySpec(in *RefinerySpec, out *kubereplay.RefinerySpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_RefinerySpec_To_kubereplay_RefinerySpec(in, out, s)
}

func autoConvert_kubereplay_RefinerySpec_To_v1alpha1_RefinerySpec(in *kubereplay.RefinerySpec, out *RefinerySpec, s conversion.Scope) error {
	out.Workers = in.Workers
	out.Timeout = in.Timeout
	out.Storage = (*RefineryStorage)(unsafe.Pointer(in.Storage))
	return nil
}

// Convert_kubereplay_RefinerySpec_To_v1alpha1_RefinerySpec is an autogenerated conversion function.
func Convert_kubereplay_RefinerySpec_To_v1alpha1_RefinerySpec(in *kubereplay.RefinerySpec, out *RefinerySpec, s conversion.Scope) error {
	return autoConvert_kubereplay_RefinerySpec_To_v1alpha1_RefinerySpec(in, out, s)
}

func autoConvert_v1alpha1_RefineryStatus_To_kubereplay_RefineryStatus(in *RefineryStatus, out *kubereplay.RefineryStatus, s conversion.Scope) error {
	out.Deployed = in.Deployed
	return nil
}

// Convert_v1alpha1_RefineryStatus_To_kubereplay_RefineryStatus is an autogenerated conversion function.
func Convert_v1alpha1_RefineryStatus_To_kubereplay_RefineryStatus(in *RefineryStatus, out *kubereplay.RefineryStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_RefineryStatus_To_kubereplay_RefineryStatus(in, out, s)
}

func autoConvert_kubereplay_RefineryStatus_To_v1alpha1_RefineryStatus(in *kubereplay.RefineryStatus, out *RefineryStatus, s conversion.Scope) error {
	out.Deployed = in.Deployed
	return nil
}

// Convert_kubereplay_RefineryStatus_To_v1alpha1_RefineryStatus is an autogenerated conversion function.
func Convert_kubereplay_RefineryStatus_To_v1alpha1_RefineryStatus(in *kubereplay.RefineryStatus, out *RefineryStatus, s conversion.Scope) error {
	return autoConvert_kubereplay_RefineryStatus_To_v1alpha1_RefineryStatus(in, out, s)
}

func autoConvert_v1alpha1_RefineryStatusStrategy_To_kubereplay_RefineryStatusStrategy(in *RefineryStatusStrategy, out *kubereplay.RefineryStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_v1alpha1_RefineryStatusStrategy_To_kubereplay_RefineryStatusStrategy is an autogenerated conversion function.
func Convert_v1alpha1_RefineryStatusStrategy_To_kubereplay_RefineryStatusStrategy(in *RefineryStatusStrategy, out *kubereplay.RefineryStatusStrategy, s conversion.Scope) error {
	return autoConvert_v1alpha1_RefineryStatusStrategy_To_kubereplay_RefineryStatusStrategy(in, out, s)
}

func autoConvert_kubereplay_RefineryStatusStrategy_To_v1alpha1_RefineryStatusStrategy(in *kubereplay.RefineryStatusStrategy, out *RefineryStatusStrategy, s conversion.Scope) error {
	out.DefaultStatusStorageStrategy = in.DefaultStatusStorageStrategy
	return nil
}

// Convert_kubereplay_RefineryStatusStrategy_To_v1alpha1_RefineryStatusStrategy is an autogenerated conversion function.
func Convert_kubereplay_RefineryStatusStrategy_To_v1alpha1_RefineryStatusStrategy(in *kubereplay.RefineryStatusStrategy, out *RefineryStatusStrategy, s conversion.Scope) error {
	return autoConvert_kubereplay_RefineryStatusStrategy_To_v1alpha1_RefineryStatusStrategy(in, out, s)
}

func autoConvert_v1alpha1_RefineryStorage_To_kubereplay_RefineryStorage(in *RefineryStorage, out *kubereplay.RefineryStorage, s conversion.Scope) error {
	out.File = (*kubereplay.FileSilo)(unsafe.Pointer(in.File))
	out.Tcp = (*kubereplay.TcpSilo)(unsafe.Pointer(in.Tcp))
	out.Stdout = (*kubereplay.StdoutSilo)(unsafe.Pointer(in.Stdout))
	out.Http = (*kubereplay.HttpSilo)(unsafe.Pointer(in.Http))
	out.Elasticsearch = (*kubereplay.ElasticsearchSilo)(unsafe.Pointer(in.Elasticsearch))
	out.Kafka = (*kubereplay.KafkaSilo)(unsafe.Pointer(in.Kafka))
	return nil
}

// Convert_v1alpha1_RefineryStorage_To_kubereplay_RefineryStorage is an autogenerated conversion function.
func Convert_v1alpha1_RefineryStorage_To_kubereplay_RefineryStorage(in *RefineryStorage, out *kubereplay.RefineryStorage, s conversion.Scope) error {
	return autoConvert_v1alpha1_RefineryStorage_To_kubereplay_RefineryStorage(in, out, s)
}

func autoConvert_kubereplay_RefineryStorage_To_v1alpha1_RefineryStorage(in *kubereplay.RefineryStorage, out *RefineryStorage, s conversion.Scope) error {
	out.File = (*FileSilo)(unsafe.Pointer(in.File))
	out.Tcp = (*TcpSilo)(unsafe.Pointer(in.Tcp))
	out.Stdout = (*StdoutSilo)(unsafe.Pointer(in.Stdout))
	out.Http = (*HttpSilo)(unsafe.Pointer(in.Http))
	out.Elasticsearch = (*ElasticsearchSilo)(unsafe.Pointer(in.Elasticsearch))
	out.Kafka = (*KafkaSilo)(unsafe.Pointer(in.Kafka))
	return nil
}

// Convert_kubereplay_RefineryStorage_To_v1alpha1_RefineryStorage is an autogenerated conversion function.
func Convert_kubereplay_RefineryStorage_To_v1alpha1_RefineryStorage(in *kubereplay.RefineryStorage, out *RefineryStorage, s conversion.Scope) error {
	return autoConvert_kubereplay_RefineryStorage_To_v1alpha1_RefineryStorage(in, out, s)
}

func autoConvert_v1alpha1_RefineryStrategy_To_kubereplay_RefineryStrategy(in *RefineryStrategy, out *kubereplay.RefineryStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_v1alpha1_RefineryStrategy_To_kubereplay_RefineryStrategy is an autogenerated conversion function.
func Convert_v1alpha1_RefineryStrategy_To_kubereplay_RefineryStrategy(in *RefineryStrategy, out *kubereplay.RefineryStrategy, s conversion.Scope) error {
	return autoConvert_v1alpha1_RefineryStrategy_To_kubereplay_RefineryStrategy(in, out, s)
}

func autoConvert_kubereplay_RefineryStrategy_To_v1alpha1_RefineryStrategy(in *kubereplay.RefineryStrategy, out *RefineryStrategy, s conversion.Scope) error {
	out.DefaultStorageStrategy = in.DefaultStorageStrategy
	return nil
}

// Convert_kubereplay_RefineryStrategy_To_v1alpha1_RefineryStrategy is an autogenerated conversion function.
func Convert_kubereplay_RefineryStrategy_To_v1alpha1_RefineryStrategy(in *kubereplay.RefineryStrategy, out *RefineryStrategy, s conversion.Scope) error {
	return autoConvert_kubereplay_RefineryStrategy_To_v1alpha1_RefineryStrategy(in, out, s)
}

func autoConvert_v1alpha1_StdoutSilo_To_kubereplay_StdoutSilo(in *StdoutSilo, out *kubereplay.StdoutSilo, s conversion.Scope) error {
	out.Enabled = in.Enabled
	return nil
}

// Convert_v1alpha1_StdoutSilo_To_kubereplay_StdoutSilo is an autogenerated conversion function.
func Convert_v1alpha1_StdoutSilo_To_kubereplay_StdoutSilo(in *StdoutSilo, out *kubereplay.StdoutSilo, s conversion.Scope) error {
	return autoConvert_v1alpha1_StdoutSilo_To_kubereplay_StdoutSilo(in, out, s)
}

func autoConvert_kubereplay_StdoutSilo_To_v1alpha1_StdoutSilo(in *kubereplay.StdoutSilo, out *StdoutSilo, s conversion.Scope) error {
	out.Enabled = in.Enabled
	return nil
}

// Convert_kubereplay_StdoutSilo_To_v1alpha1_StdoutSilo is an autogenerated conversion function.
func Convert_kubereplay_StdoutSilo_To_v1alpha1_StdoutSilo(in *kubereplay.StdoutSilo, out *StdoutSilo, s conversion.Scope) error {
	return autoConvert_kubereplay_StdoutSilo_To_v1alpha1_StdoutSilo(in, out, s)
}

func autoConvert_v1alpha1_TcpSilo_To_kubereplay_TcpSilo(in *TcpSilo, out *kubereplay.TcpSilo, s conversion.Scope) error {
	out.Enabled = in.Enabled
	out.Uri = in.Uri
	return nil
}

// Convert_v1alpha1_TcpSilo_To_kubereplay_TcpSilo is an autogenerated conversion function.
func Convert_v1alpha1_TcpSilo_To_kubereplay_TcpSilo(in *TcpSilo, out *kubereplay.TcpSilo, s conversion.Scope) error {
	return autoConvert_v1alpha1_TcpSilo_To_kubereplay_TcpSilo(in, out, s)
}

func autoConvert_kubereplay_TcpSilo_To_v1alpha1_TcpSilo(in *kubereplay.TcpSilo, out *TcpSilo, s conversion.Scope) error {
	out.Enabled = in.Enabled
	out.Uri = in.Uri
	return nil
}

// Convert_kubereplay_TcpSilo_To_v1alpha1_TcpSilo is an autogenerated conversion function.
func Convert_kubereplay_TcpSilo_To_v1alpha1_TcpSilo(in *kubereplay.TcpSilo, out *TcpSilo, s conversion.Scope) error {
	return autoConvert_kubereplay_TcpSilo_To_v1alpha1_TcpSilo(in, out, s)
}
