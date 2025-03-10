// Code generated by applyconfiguration-gen. DO NOT EDIT.

package internal

import (
	fmt "fmt"
	sync "sync"

	typed "sigs.k8s.io/structured-merge-diff/v4/typed"
)

func Parser() *typed.Parser {
	parserOnce.Do(func() {
		var err error
		parser, err = typed.NewParser(schemaYAML)
		if err != nil {
			panic(fmt.Sprintf("Failed to parse schema: %v", err))
		}
	})
	return parser
}

var parserOnce sync.Once
var parser *typed.Parser
var schemaYAML = typed.YAMLObject(`types:
- name: com.github.openshift.api.authorization.v1.ClusterRole
  map:
    fields:
    - name: aggregationRule
      type:
        namedType: io.k8s.api.rbac.v1.AggregationRule
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: rules
      type:
        list:
          elementType:
            namedType: com.github.openshift.api.authorization.v1.PolicyRule
          elementRelationship: atomic
- name: com.github.openshift.api.authorization.v1.ClusterRoleBinding
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: groupNames
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: roleRef
      type:
        namedType: io.k8s.api.core.v1.ObjectReference
      default: {}
    - name: subjects
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.ObjectReference
          elementRelationship: atomic
    - name: userNames
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
- name: com.github.openshift.api.authorization.v1.GroupRestriction
  map:
    fields:
    - name: groups
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: labels
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
          elementRelationship: atomic
- name: com.github.openshift.api.authorization.v1.PolicyRule
  map:
    fields:
    - name: apiGroups
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: attributeRestrictions
      type:
        namedType: __untyped_atomic_
    - name: nonResourceURLs
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: resourceNames
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: resources
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: verbs
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
- name: com.github.openshift.api.authorization.v1.Role
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: rules
      type:
        list:
          elementType:
            namedType: com.github.openshift.api.authorization.v1.PolicyRule
          elementRelationship: atomic
- name: com.github.openshift.api.authorization.v1.RoleBinding
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: groupNames
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: roleRef
      type:
        namedType: io.k8s.api.core.v1.ObjectReference
      default: {}
    - name: subjects
      type:
        list:
          elementType:
            namedType: io.k8s.api.core.v1.ObjectReference
          elementRelationship: atomic
    - name: userNames
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
- name: com.github.openshift.api.authorization.v1.RoleBindingRestriction
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: metadata
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
      default: {}
    - name: spec
      type:
        namedType: com.github.openshift.api.authorization.v1.RoleBindingRestrictionSpec
      default: {}
- name: com.github.openshift.api.authorization.v1.RoleBindingRestrictionSpec
  map:
    fields:
    - name: grouprestriction
      type:
        namedType: com.github.openshift.api.authorization.v1.GroupRestriction
    - name: serviceaccountrestriction
      type:
        namedType: com.github.openshift.api.authorization.v1.ServiceAccountRestriction
    - name: userrestriction
      type:
        namedType: com.github.openshift.api.authorization.v1.UserRestriction
- name: com.github.openshift.api.authorization.v1.ServiceAccountReference
  map:
    fields:
    - name: name
      type:
        scalar: string
      default: ""
    - name: namespace
      type:
        scalar: string
      default: ""
- name: com.github.openshift.api.authorization.v1.ServiceAccountRestriction
  map:
    fields:
    - name: namespaces
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: serviceaccounts
      type:
        list:
          elementType:
            namedType: com.github.openshift.api.authorization.v1.ServiceAccountReference
          elementRelationship: atomic
- name: com.github.openshift.api.authorization.v1.UserRestriction
  map:
    fields:
    - name: groups
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
    - name: labels
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
          elementRelationship: atomic
    - name: users
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
- name: io.k8s.api.core.v1.ObjectReference
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: fieldPath
      type:
        scalar: string
    - name: kind
      type:
        scalar: string
    - name: name
      type:
        scalar: string
    - name: namespace
      type:
        scalar: string
    - name: resourceVersion
      type:
        scalar: string
    - name: uid
      type:
        scalar: string
    elementRelationship: atomic
- name: io.k8s.api.rbac.v1.AggregationRule
  map:
    fields:
    - name: clusterRoleSelectors
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
          elementRelationship: atomic
- name: io.k8s.apimachinery.pkg.apis.meta.v1.FieldsV1
  map:
    elementType:
      scalar: untyped
      list:
        elementType:
          namedType: __untyped_atomic_
        elementRelationship: atomic
      map:
        elementType:
          namedType: __untyped_deduced_
        elementRelationship: separable
- name: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelector
  map:
    fields:
    - name: matchExpressions
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelectorRequirement
          elementRelationship: atomic
    - name: matchLabels
      type:
        map:
          elementType:
            scalar: string
    elementRelationship: atomic
- name: io.k8s.apimachinery.pkg.apis.meta.v1.LabelSelectorRequirement
  map:
    fields:
    - name: key
      type:
        scalar: string
      default: ""
    - name: operator
      type:
        scalar: string
      default: ""
    - name: values
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: atomic
- name: io.k8s.apimachinery.pkg.apis.meta.v1.ManagedFieldsEntry
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
    - name: fieldsType
      type:
        scalar: string
    - name: fieldsV1
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.FieldsV1
    - name: manager
      type:
        scalar: string
    - name: operation
      type:
        scalar: string
    - name: subresource
      type:
        scalar: string
    - name: time
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
- name: io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta
  map:
    fields:
    - name: annotations
      type:
        map:
          elementType:
            scalar: string
    - name: creationTimestamp
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
    - name: deletionGracePeriodSeconds
      type:
        scalar: numeric
    - name: deletionTimestamp
      type:
        namedType: io.k8s.apimachinery.pkg.apis.meta.v1.Time
    - name: finalizers
      type:
        list:
          elementType:
            scalar: string
          elementRelationship: associative
    - name: generateName
      type:
        scalar: string
    - name: generation
      type:
        scalar: numeric
    - name: labels
      type:
        map:
          elementType:
            scalar: string
    - name: managedFields
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.ManagedFieldsEntry
          elementRelationship: atomic
    - name: name
      type:
        scalar: string
    - name: namespace
      type:
        scalar: string
    - name: ownerReferences
      type:
        list:
          elementType:
            namedType: io.k8s.apimachinery.pkg.apis.meta.v1.OwnerReference
          elementRelationship: associative
          keys:
          - uid
    - name: resourceVersion
      type:
        scalar: string
    - name: selfLink
      type:
        scalar: string
    - name: uid
      type:
        scalar: string
- name: io.k8s.apimachinery.pkg.apis.meta.v1.OwnerReference
  map:
    fields:
    - name: apiVersion
      type:
        scalar: string
      default: ""
    - name: blockOwnerDeletion
      type:
        scalar: boolean
    - name: controller
      type:
        scalar: boolean
    - name: kind
      type:
        scalar: string
      default: ""
    - name: name
      type:
        scalar: string
      default: ""
    - name: uid
      type:
        scalar: string
      default: ""
    elementRelationship: atomic
- name: io.k8s.apimachinery.pkg.apis.meta.v1.Time
  scalar: untyped
- name: io.k8s.apimachinery.pkg.runtime.RawExtension
  map:
    elementType:
      scalar: untyped
      list:
        elementType:
          namedType: __untyped_atomic_
        elementRelationship: atomic
      map:
        elementType:
          namedType: __untyped_deduced_
        elementRelationship: separable
- name: __untyped_atomic_
  scalar: untyped
  list:
    elementType:
      namedType: __untyped_atomic_
    elementRelationship: atomic
  map:
    elementType:
      namedType: __untyped_atomic_
    elementRelationship: atomic
- name: __untyped_deduced_
  scalar: untyped
  list:
    elementType:
      namedType: __untyped_atomic_
    elementRelationship: atomic
  map:
    elementType:
      namedType: __untyped_deduced_
    elementRelationship: separable
`)
