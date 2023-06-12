// Code generated by ogen, DO NOT EDIT.

package api

import (
	"github.com/go-faster/jx"

	std "encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAnyTest_EncodeDecode(t *testing.T) {
	var typ AnyTest
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 AnyTest
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestAnyTestAnyMap_EncodeDecode(t *testing.T) {
	var typ AnyTestAnyMap
	typ = make(AnyTestAnyMap)
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 AnyTestAnyMap
	typ2 = make(AnyTestAnyMap)
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestArrayTest_EncodeDecode(t *testing.T) {
	var typ ArrayTest
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 ArrayTest
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestData_EncodeDecode(t *testing.T) {
	var typ Data
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Data
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestDataDescription_EncodeDecode(t *testing.T) {
	var typ DataDescription
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 DataDescription
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestDefaultTest_EncodeDecode(t *testing.T) {
	var typ DefaultTest
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 DefaultTest
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestDefaultTestEnum_EncodeDecode(t *testing.T) {
	var typ DefaultTestEnum
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 DefaultTestEnum
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestDescriptionDetailed_EncodeDecode(t *testing.T) {
	var typ DescriptionDetailed
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 DescriptionDetailed
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestDescriptionSimple_EncodeDecode(t *testing.T) {
	var typ DescriptionSimple
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 DescriptionSimple
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestError_EncodeDecode(t *testing.T) {
	var typ Error
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Error
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestID_EncodeDecode(t *testing.T) {
	var typ ID
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 ID
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestIssue143_EncodeDecode(t *testing.T) {
	var typ Issue143
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Issue143
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestIssue1430_EncodeDecode(t *testing.T) {
	var typ Issue1430
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Issue1430
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestIssue1431_EncodeDecode(t *testing.T) {
	var typ Issue1431
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Issue1431
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestIssue1432_EncodeDecode(t *testing.T) {
	var typ Issue1432
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Issue1432
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestIssue1433_EncodeDecode(t *testing.T) {
	var typ Issue1433
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Issue1433
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestIssue943_EncodeDecode(t *testing.T) {
	var typ Issue943
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Issue943
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestIssue943Map_EncodeDecode(t *testing.T) {
	var typ Issue943Map
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Issue943Map
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestIssue943MapPattern0_EncodeDecode(t *testing.T) {
	var typ Issue943MapPattern0
	typ = make(Issue943MapPattern0)
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Issue943MapPattern0
	typ2 = make(Issue943MapPattern0)
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestIssue943Variant1_EncodeDecode(t *testing.T) {
	var typ Issue943Variant1
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Issue943Variant1
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestIssue943Variant2_EncodeDecode(t *testing.T) {
	var typ Issue943Variant2
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Issue943Variant2
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestMapWithProperties_EncodeDecode(t *testing.T) {
	var typ MapWithProperties
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 MapWithProperties
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestMapWithPropertiesAdditional_EncodeDecode(t *testing.T) {
	var typ MapWithPropertiesAdditional
	typ = make(MapWithPropertiesAdditional)
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 MapWithPropertiesAdditional
	typ2 = make(MapWithPropertiesAdditional)
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestMapWithPropertiesInlinedSubMap_EncodeDecode(t *testing.T) {
	var typ MapWithPropertiesInlinedSubMap
	typ = make(MapWithPropertiesInlinedSubMap)
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 MapWithPropertiesInlinedSubMap
	typ2 = make(MapWithPropertiesInlinedSubMap)
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestMaxPropertiesTest_EncodeDecode(t *testing.T) {
	var typ MaxPropertiesTest
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 MaxPropertiesTest
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestNoAdditionalPropertiesTest_EncodeDecode(t *testing.T) {
	var typ NoAdditionalPropertiesTest
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 NoAdditionalPropertiesTest
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestNullValue_EncodeDecode(t *testing.T) {
	var typ NullValue
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 NullValue
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestNullableEnums_EncodeDecode(t *testing.T) {
	var typ NullableEnums
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 NullableEnums
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestNullableEnumsBoth_EncodeDecode(t *testing.T) {
	var typ NullableEnumsBoth
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 NullableEnumsBoth
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestNullableEnumsOnlyNullValue_EncodeDecode(t *testing.T) {
	var typ NullableEnumsOnlyNullValue
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 NullableEnumsOnlyNullValue
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestNullableEnumsOnlyNullable_EncodeDecode(t *testing.T) {
	var typ NullableEnumsOnlyNullable
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 NullableEnumsOnlyNullable
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestOneOfBooleanSumNullables_EncodeDecode(t *testing.T) {
	var typ OneOfBooleanSumNullables
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 OneOfBooleanSumNullables
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestOneOfBugs_EncodeDecode(t *testing.T) {
	var typ OneOfBugs
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 OneOfBugs
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestOneOfMappingReference_EncodeDecode(t *testing.T) {
	var typ OneOfMappingReference
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 OneOfMappingReference
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestOneOfMappingReferenceA_EncodeDecode(t *testing.T) {
	var typ OneOfMappingReferenceA
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 OneOfMappingReferenceA
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestOneOfMappingReferenceB_EncodeDecode(t *testing.T) {
	var typ OneOfMappingReferenceB
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 OneOfMappingReferenceB
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestOneOfMappingReferenceBData_EncodeDecode(t *testing.T) {
	var typ OneOfMappingReferenceBData
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 OneOfMappingReferenceBData
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestOneOfNullables_EncodeDecode(t *testing.T) {
	var typ OneOfNullables
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 OneOfNullables
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestOneOfUUIDAndIntEnum_EncodeDecode(t *testing.T) {
	var typ OneOfUUIDAndIntEnum
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 OneOfUUIDAndIntEnum
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestOneOfUUIDAndIntEnum1_EncodeDecode(t *testing.T) {
	var typ OneOfUUIDAndIntEnum1
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 OneOfUUIDAndIntEnum1
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestOneOfWithNullable_EncodeDecode(t *testing.T) {
	var typ OneOfWithNullable
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 OneOfWithNullable
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestOnePropertyObject_EncodeDecode(t *testing.T) {
	var typ OnePropertyObject
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 OnePropertyObject
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestOneVariantHasNoUniqueFields_EncodeDecode(t *testing.T) {
	var typ OneVariantHasNoUniqueFields
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 OneVariantHasNoUniqueFields
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestOneVariantHasNoUniqueFields0_EncodeDecode(t *testing.T) {
	var typ OneVariantHasNoUniqueFields0
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 OneVariantHasNoUniqueFields0
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestOneVariantHasNoUniqueFields1_EncodeDecode(t *testing.T) {
	var typ OneVariantHasNoUniqueFields1
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 OneVariantHasNoUniqueFields1
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestOnlyEmptyObject_EncodeDecode(t *testing.T) {
	var typ OnlyEmptyObject
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 OnlyEmptyObject
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestOnlyPatternedPropsObject_EncodeDecode(t *testing.T) {
	var typ OnlyPatternedPropsObject
	typ = make(OnlyPatternedPropsObject)
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 OnlyPatternedPropsObject
	typ2 = make(OnlyPatternedPropsObject)
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestPatternRecursiveMap_EncodeDecode(t *testing.T) {
	var typ PatternRecursiveMap
	typ = make(PatternRecursiveMap)
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 PatternRecursiveMap
	typ2 = make(PatternRecursiveMap)
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestPet_EncodeDecode(t *testing.T) {
	var typ Pet
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Pet
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestPetGetDef_EncodeDecode(t *testing.T) {
	var typ PetGetDef
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 PetGetDef
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestPetKind_EncodeDecode(t *testing.T) {
	var typ PetKind
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 PetKind
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestPetName_EncodeDecode(t *testing.T) {
	var typ PetName
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 PetName
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestPetType_EncodeDecode(t *testing.T) {
	var typ PetType
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 PetType
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestRecursiveArray_EncodeDecode(t *testing.T) {
	var typ RecursiveArray
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 RecursiveArray
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestRecursiveMap_EncodeDecode(t *testing.T) {
	var typ RecursiveMap
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 RecursiveMap
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestRecursiveMapAdditional_EncodeDecode(t *testing.T) {
	var typ RecursiveMapAdditional
	typ = make(RecursiveMapAdditional)
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 RecursiveMapAdditional
	typ2 = make(RecursiveMapAdditional)
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestStringIntMap_EncodeDecode(t *testing.T) {
	var typ StringIntMap
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 StringIntMap
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestStringIntMapAdditional_EncodeDecode(t *testing.T) {
	var typ StringIntMapAdditional
	typ = make(StringIntMapAdditional)
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 StringIntMapAdditional
	typ2 = make(StringIntMapAdditional)
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestStringIntMapPattern0_EncodeDecode(t *testing.T) {
	var typ StringIntMapPattern0
	typ = make(StringIntMapPattern0)
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 StringIntMapPattern0
	typ2 = make(StringIntMapPattern0)
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestStringMap_EncodeDecode(t *testing.T) {
	var typ StringMap
	typ = make(StringMap)
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 StringMap
	typ2 = make(StringMap)
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestStringStringMap_EncodeDecode(t *testing.T) {
	var typ StringStringMap
	typ = make(StringStringMap)
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 StringStringMap
	typ2 = make(StringStringMap)
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestTestFloatValidation_EncodeDecode(t *testing.T) {
	var typ TestFloatValidation
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 TestFloatValidation
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestTestNullableOneofsCreated_EncodeDecode(t *testing.T) {
	var typ TestNullableOneofsCreated
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 TestNullableOneofsCreated
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestTestNullableOneofsOK_EncodeDecode(t *testing.T) {
	var typ TestNullableOneofsOK
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 TestNullableOneofsOK
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestTupleTest_EncodeDecode(t *testing.T) {
	var typ TupleTest
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 TupleTest
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestTupleTestV4_EncodeDecode(t *testing.T) {
	var typ TupleTestV4
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 TupleTestV4
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestUniqueItemsTest_EncodeDecode(t *testing.T) {
	var typ UniqueItemsTest
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 UniqueItemsTest
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestValidationStringMap_EncodeDecode(t *testing.T) {
	var typ ValidationStringMap
	typ = make(ValidationStringMap)
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 ValidationStringMap
	typ2 = make(ValidationStringMap)
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
