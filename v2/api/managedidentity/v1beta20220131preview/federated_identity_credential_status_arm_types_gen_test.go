// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20220131preview

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_FederatedIdentityCredential_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FederatedIdentityCredential_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFederatedIdentityCredential_STATUSARM, FederatedIdentityCredential_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFederatedIdentityCredential_STATUSARM runs a test to see if a specific instance of FederatedIdentityCredential_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFederatedIdentityCredential_STATUSARM(subject FederatedIdentityCredential_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FederatedIdentityCredential_STATUSARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of FederatedIdentityCredential_STATUSARM instances for property testing - lazily instantiated by
// FederatedIdentityCredential_STATUSARMGenerator()
var federatedIdentityCredential_STATUSARMGenerator gopter.Gen

// FederatedIdentityCredential_STATUSARMGenerator returns a generator of FederatedIdentityCredential_STATUSARM instances for property testing.
// We first initialize federatedIdentityCredential_STATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FederatedIdentityCredential_STATUSARMGenerator() gopter.Gen {
	if federatedIdentityCredential_STATUSARMGenerator != nil {
		return federatedIdentityCredential_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFederatedIdentityCredential_STATUSARM(generators)
	federatedIdentityCredential_STATUSARMGenerator = gen.Struct(reflect.TypeOf(FederatedIdentityCredential_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFederatedIdentityCredential_STATUSARM(generators)
	AddRelatedPropertyGeneratorsForFederatedIdentityCredential_STATUSARM(generators)
	federatedIdentityCredential_STATUSARMGenerator = gen.Struct(reflect.TypeOf(FederatedIdentityCredential_STATUSARM{}), generators)

	return federatedIdentityCredential_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForFederatedIdentityCredential_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFederatedIdentityCredential_STATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFederatedIdentityCredential_STATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFederatedIdentityCredential_STATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(FederatedIdentityCredentialProperties_STATUSARMGenerator())
}

func Test_FederatedIdentityCredentialProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FederatedIdentityCredentialProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFederatedIdentityCredentialProperties_STATUSARM, FederatedIdentityCredentialProperties_STATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFederatedIdentityCredentialProperties_STATUSARM runs a test to see if a specific instance of FederatedIdentityCredentialProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFederatedIdentityCredentialProperties_STATUSARM(subject FederatedIdentityCredentialProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FederatedIdentityCredentialProperties_STATUSARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of FederatedIdentityCredentialProperties_STATUSARM instances for property testing - lazily instantiated by
// FederatedIdentityCredentialProperties_STATUSARMGenerator()
var federatedIdentityCredentialProperties_STATUSARMGenerator gopter.Gen

// FederatedIdentityCredentialProperties_STATUSARMGenerator returns a generator of FederatedIdentityCredentialProperties_STATUSARM instances for property testing.
func FederatedIdentityCredentialProperties_STATUSARMGenerator() gopter.Gen {
	if federatedIdentityCredentialProperties_STATUSARMGenerator != nil {
		return federatedIdentityCredentialProperties_STATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFederatedIdentityCredentialProperties_STATUSARM(generators)
	federatedIdentityCredentialProperties_STATUSARMGenerator = gen.Struct(reflect.TypeOf(FederatedIdentityCredentialProperties_STATUSARM{}), generators)

	return federatedIdentityCredentialProperties_STATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForFederatedIdentityCredentialProperties_STATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFederatedIdentityCredentialProperties_STATUSARM(gens map[string]gopter.Gen) {
	gens["Audiences"] = gen.SliceOf(gen.AlphaString())
	gens["Issuer"] = gen.PtrOf(gen.AlphaString())
	gens["Subject"] = gen.PtrOf(gen.AlphaString())
}
