// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v2beta2 "k8s.io/api/autoscaling/v2beta2"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Algorithm) DeepCopyInto(out *Algorithm) {
	*out = *in
	if in.DSP != nil {
		in, out := &in.DSP, &out.DSP
		*out = new(Dsp)
		(*in).DeepCopyInto(*out)
	}
	if in.Percentile != nil {
		in, out := &in.Percentile, &out.Percentile
		*out = new(Percentile)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Algorithm.
func (in *Algorithm) DeepCopy() *Algorithm {
	if in == nil {
		return nil
	}
	out := new(Algorithm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Dsp) DeepCopyInto(out *Dsp) {
	*out = *in
	in.Estimator.DeepCopyInto(&out.Estimator)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Dsp.
func (in *Dsp) DeepCopy() *Dsp {
	if in == nil {
		return nil
	}
	out := new(Dsp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Estimator) DeepCopyInto(out *Estimator) {
	*out = *in
	if in.MaxValueEstimators != nil {
		in, out := &in.MaxValueEstimators, &out.MaxValueEstimators
		*out = make([]*MaxValueEstimator, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(MaxValueEstimator)
				**out = **in
			}
		}
	}
	if in.FFTEstimators != nil {
		in, out := &in.FFTEstimators, &out.FFTEstimators
		*out = make([]*FFTEstimator, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(FFTEstimator)
				**out = **in
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Estimator.
func (in *Estimator) DeepCopy() *Estimator {
	if in == nil {
		return nil
	}
	out := new(Estimator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FFTEstimator) DeepCopyInto(out *FFTEstimator) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FFTEstimator.
func (in *FFTEstimator) DeepCopy() *FFTEstimator {
	if in == nil {
		return nil
	}
	out := new(FFTEstimator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HistogramConfig) DeepCopyInto(out *HistogramConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HistogramConfig.
func (in *HistogramConfig) DeepCopy() *HistogramConfig {
	if in == nil {
		return nil
	}
	out := new(HistogramConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Label) DeepCopyInto(out *Label) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Label.
func (in *Label) DeepCopy() *Label {
	if in == nil {
		return nil
	}
	out := new(Label)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaxValueEstimator) DeepCopyInto(out *MaxValueEstimator) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaxValueEstimator.
func (in *MaxValueEstimator) DeepCopy() *MaxValueEstimator {
	if in == nil {
		return nil
	}
	out := new(MaxValueEstimator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricPredictionConfig) DeepCopyInto(out *MetricPredictionConfig) {
	*out = *in
	if in.DSP != nil {
		in, out := &in.DSP, &out.DSP
		*out = new(Dsp)
		(*in).DeepCopyInto(*out)
	}
	if in.Percentile != nil {
		in, out := &in.Percentile, &out.Percentile
		*out = new(Percentile)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricPredictionConfig.
func (in *MetricPredictionConfig) DeepCopy() *MetricPredictionConfig {
	if in == nil {
		return nil
	}
	out := new(MetricPredictionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricSelector) DeepCopyInto(out *MetricSelector) {
	*out = *in
	if in.QueryConditions != nil {
		in, out := &in.QueryConditions, &out.QueryConditions
		*out = make([]QueryCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricSelector.
func (in *MetricSelector) DeepCopy() *MetricSelector {
	if in == nil {
		return nil
	}
	out := new(MetricSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricTimeSeries) DeepCopyInto(out *MetricTimeSeries) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]Label, len(*in))
		copy(*out, *in)
	}
	if in.Samples != nil {
		in, out := &in.Samples, &out.Samples
		*out = make([]Sample, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricTimeSeries.
func (in *MetricTimeSeries) DeepCopy() *MetricTimeSeries {
	if in == nil {
		return nil
	}
	out := new(MetricTimeSeries)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in MetricTimeSeriesList) DeepCopyInto(out *MetricTimeSeriesList) {
	{
		in := &in
		*out = make(MetricTimeSeriesList, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(MetricTimeSeries)
				(*in).DeepCopyInto(*out)
			}
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricTimeSeriesList.
func (in MetricTimeSeriesList) DeepCopy() MetricTimeSeriesList {
	if in == nil {
		return nil
	}
	out := new(MetricTimeSeriesList)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePrediction) DeepCopyInto(out *NodePrediction) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePrediction.
func (in *NodePrediction) DeepCopy() *NodePrediction {
	if in == nil {
		return nil
	}
	out := new(NodePrediction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodePrediction) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePredictionList) DeepCopyInto(out *NodePredictionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodePrediction, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePredictionList.
func (in *NodePredictionList) DeepCopy() *NodePredictionList {
	if in == nil {
		return nil
	}
	out := new(NodePredictionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodePredictionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePredictionResourceSpec) DeepCopyInto(out *NodePredictionResourceSpec) {
	*out = *in
	out.Period = in.Period
	if in.MetricPredictionConfigs != nil {
		in, out := &in.MetricPredictionConfigs, &out.MetricPredictionConfigs
		*out = make([]MetricPredictionConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePredictionResourceSpec.
func (in *NodePredictionResourceSpec) DeepCopy() *NodePredictionResourceSpec {
	if in == nil {
		return nil
	}
	out := new(NodePredictionResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePredictionResourceStatus) DeepCopyInto(out *NodePredictionResourceStatus) {
	*out = *in
	if in.NextPossible != nil {
		in, out := &in.NextPossible, &out.NextPossible
		*out = make(Prediction, len(*in))
		for key, val := range *in {
			var outVal []*Vector
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(TimeSeries, len(*in))
				for i := range *in {
					if (*in)[i] != nil {
						in, out := &(*in)[i], &(*out)[i]
						*out = new(Vector)
						**out = **in
					}
				}
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePredictionResourceStatus.
func (in *NodePredictionResourceStatus) DeepCopy() *NodePredictionResourceStatus {
	if in == nil {
		return nil
	}
	out := new(NodePredictionResourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Percentile) DeepCopyInto(out *Percentile) {
	*out = *in
	out.Histogram = in.Histogram
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Percentile.
func (in *Percentile) DeepCopy() *Percentile {
	if in == nil {
		return nil
	}
	out := new(Percentile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodGroupPrediction) DeepCopyInto(out *PodGroupPrediction) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodGroupPrediction.
func (in *PodGroupPrediction) DeepCopy() *PodGroupPrediction {
	if in == nil {
		return nil
	}
	out := new(PodGroupPrediction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodGroupPrediction) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodGroupPredictionCondition) DeepCopyInto(out *PodGroupPredictionCondition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodGroupPredictionCondition.
func (in *PodGroupPredictionCondition) DeepCopy() *PodGroupPredictionCondition {
	if in == nil {
		return nil
	}
	out := new(PodGroupPredictionCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodGroupPredictionList) DeepCopyInto(out *PodGroupPredictionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PodGroupPrediction, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodGroupPredictionList.
func (in *PodGroupPredictionList) DeepCopy() *PodGroupPredictionList {
	if in == nil {
		return nil
	}
	out := new(PodGroupPredictionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodGroupPredictionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodGroupPredictionSpec) DeepCopyInto(out *PodGroupPredictionSpec) {
	*out = *in
	if in.Start != nil {
		in, out := &in.Start, &out.Start
		*out = (*in).DeepCopy()
	}
	if in.End != nil {
		in, out := &in.End, &out.End
		*out = (*in).DeepCopy()
	}
	out.PredictionWindow = in.PredictionWindow
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.WorkloadRef != nil {
		in, out := &in.WorkloadRef, &out.WorkloadRef
		*out = new(v2beta2.CrossVersionObjectReference)
		**out = **in
	}
	in.LabelSelector.DeepCopyInto(&out.LabelSelector)
	if in.MetricPredictionConfigs != nil {
		in, out := &in.MetricPredictionConfigs, &out.MetricPredictionConfigs
		*out = make([]MetricPredictionConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodGroupPredictionSpec.
func (in *PodGroupPredictionSpec) DeepCopy() *PodGroupPredictionSpec {
	if in == nil {
		return nil
	}
	out := new(PodGroupPredictionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodGroupPredictionStatus) DeepCopyInto(out *PodGroupPredictionStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]PodGroupPredictionCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Aggregation != nil {
		in, out := &in.Aggregation, &out.Aggregation
		*out = make(Prediction, len(*in))
		for key, val := range *in {
			var outVal []*Vector
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(TimeSeries, len(*in))
				for i := range *in {
					if (*in)[i] != nil {
						in, out := &(*in)[i], &(*out)[i]
						*out = new(Vector)
						**out = **in
					}
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make(map[string]Prediction, len(*in))
		for key, val := range *in {
			var outVal map[string]TimeSeries
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(Prediction, len(*in))
				for key, val := range *in {
					var outVal []*Vector
					if val == nil {
						(*out)[key] = nil
					} else {
						in, out := &val, &outVal
						*out = make(TimeSeries, len(*in))
						for i := range *in {
							if (*in)[i] != nil {
								in, out := &(*in)[i], &(*out)[i]
								*out = new(Vector)
								**out = **in
							}
						}
					}
					(*out)[key] = outVal
				}
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodGroupPredictionStatus.
func (in *PodGroupPredictionStatus) DeepCopy() *PodGroupPredictionStatus {
	if in == nil {
		return nil
	}
	out := new(PodGroupPredictionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Prediction) DeepCopyInto(out *Prediction) {
	{
		in := &in
		*out = make(Prediction, len(*in))
		for key, val := range *in {
			var outVal []*Vector
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(TimeSeries, len(*in))
				for i := range *in {
					if (*in)[i] != nil {
						in, out := &(*in)[i], &(*out)[i]
						*out = new(Vector)
						**out = **in
					}
				}
			}
			(*out)[key] = outVal
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Prediction.
func (in Prediction) DeepCopy() Prediction {
	if in == nil {
		return nil
	}
	out := new(Prediction)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredictionMetric) DeepCopyInto(out *PredictionMetric) {
	*out = *in
	if in.MetricSelector != nil {
		in, out := &in.MetricSelector, &out.MetricSelector
		*out = new(MetricSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Query != nil {
		in, out := &in.Query, &out.Query
		*out = new(Query)
		**out = **in
	}
	in.Algorithm.DeepCopyInto(&out.Algorithm)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredictionMetric.
func (in *PredictionMetric) DeepCopy() *PredictionMetric {
	if in == nil {
		return nil
	}
	out := new(PredictionMetric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Query) DeepCopyInto(out *Query) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Query.
func (in *Query) DeepCopy() *Query {
	if in == nil {
		return nil
	}
	out := new(Query)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueryCondition) DeepCopyInto(out *QueryCondition) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueryCondition.
func (in *QueryCondition) DeepCopy() *QueryCondition {
	if in == nil {
		return nil
	}
	out := new(QueryCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sample) DeepCopyInto(out *Sample) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sample.
func (in *Sample) DeepCopy() *Sample {
	if in == nil {
		return nil
	}
	out := new(Sample)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in TimeSeries) DeepCopyInto(out *TimeSeries) {
	{
		in := &in
		*out = make(TimeSeries, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Vector)
				**out = **in
			}
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeSeries.
func (in TimeSeries) DeepCopy() TimeSeries {
	if in == nil {
		return nil
	}
	out := new(TimeSeries)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeSeriesPrediction) DeepCopyInto(out *TimeSeriesPrediction) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeSeriesPrediction.
func (in *TimeSeriesPrediction) DeepCopy() *TimeSeriesPrediction {
	if in == nil {
		return nil
	}
	out := new(TimeSeriesPrediction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TimeSeriesPrediction) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeSeriesPredictionCondition) DeepCopyInto(out *TimeSeriesPredictionCondition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeSeriesPredictionCondition.
func (in *TimeSeriesPredictionCondition) DeepCopy() *TimeSeriesPredictionCondition {
	if in == nil {
		return nil
	}
	out := new(TimeSeriesPredictionCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeSeriesPredictionList) DeepCopyInto(out *TimeSeriesPredictionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TimeSeriesPrediction, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeSeriesPredictionList.
func (in *TimeSeriesPredictionList) DeepCopy() *TimeSeriesPredictionList {
	if in == nil {
		return nil
	}
	out := new(TimeSeriesPredictionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TimeSeriesPredictionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeSeriesPredictionSpec) DeepCopyInto(out *TimeSeriesPredictionSpec) {
	*out = *in
	if in.PredictionMetrics != nil {
		in, out := &in.PredictionMetrics, &out.PredictionMetrics
		*out = make([]PredictionMetric, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeSeriesPredictionSpec.
func (in *TimeSeriesPredictionSpec) DeepCopy() *TimeSeriesPredictionSpec {
	if in == nil {
		return nil
	}
	out := new(TimeSeriesPredictionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeSeriesPredictionStatus) DeepCopyInto(out *TimeSeriesPredictionStatus) {
	*out = *in
	if in.PredictionMetrics != nil {
		in, out := &in.PredictionMetrics, &out.PredictionMetrics
		*out = make(map[string]MetricTimeSeriesList, len(*in))
		for key, val := range *in {
			var outVal []*MetricTimeSeries
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(MetricTimeSeriesList, len(*in))
				for i := range *in {
					if (*in)[i] != nil {
						in, out := &(*in)[i], &(*out)[i]
						*out = new(MetricTimeSeries)
						(*in).DeepCopyInto(*out)
					}
				}
			}
			(*out)[key] = outVal
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]TimeSeriesPredictionCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeSeriesPredictionStatus.
func (in *TimeSeriesPredictionStatus) DeepCopy() *TimeSeriesPredictionStatus {
	if in == nil {
		return nil
	}
	out := new(TimeSeriesPredictionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Vector) DeepCopyInto(out *Vector) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Vector.
func (in *Vector) DeepCopy() *Vector {
	if in == nil {
		return nil
	}
	out := new(Vector)
	in.DeepCopyInto(out)
	return out
}
