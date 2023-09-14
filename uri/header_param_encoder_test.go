package uri

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestHeaderParamEncoder(t *testing.T) {
	t.Run("primitive", func(t *testing.T) {
		tests := []struct {
			HeaderName string
			Input      string
			Explode    bool
			Expect     http.Header
		}{
			{
				HeaderName: "X-MyHeader",
				Input:      "5",
				Explode:    false,
				Expect: http.Header{
					"X-Myheader": []string{"5"},
				},
			},
			{
				HeaderName: "X-MyHeader",
				Input:      "5",
				Explode:    true,
				Expect: http.Header{
					"X-Myheader": []string{"5"},
				},
			},
		}
		for i, test := range tests {
			e := headerParamEncoder{
				receiver:  newReceiver(),
				paramName: test.HeaderName,
				explode:   test.Explode,
				header:    make(http.Header),
			}
			require.NoError(t, e.EncodeValue(test.Input))
			require.NoError(t, e.serialize())

			require.Equal(t, test.Expect, e.header, fmt.Sprintf("Test %d", i+1))
		}
	})

	t.Run("array", func(t *testing.T) {
		tests := []struct {
			HeaderName string
			Input      []string
			Explode    bool
			Expect     http.Header
		}{
			{
				HeaderName: "X-MyHeader",
				Input:      []string{"3", "4", "5"},
				Explode:    false,
				Expect: http.Header{
					"X-Myheader": []string{"3,4,5"},
				},
			},
			{
				HeaderName: "X-MyHeader",
				Input:      []string{"3", "4", "5"},
				Explode:    true,
				Expect: http.Header{
					"X-Myheader": []string{"3,4,5"},
				},
			},
			{
				HeaderName: "Set-Cookie",
				Input: []string{
					(&http.Cookie{
						Name:  "key1",
						Value: "value1",
					}).String(),
					(&http.Cookie{
						Name:    "key2",
						Value:   "value2",
						Expires: time.Date(2023, 9, 6, 10, 11, 12, 0, time.UTC),
					}).String(),
				},
				Explode: true,
				Expect: http.Header{
					"Set-Cookie": []string{
						"key1=value1",
						"key2=value2; Expires=Wed, 06 Sep 2023 10:11:12 GMT",
					},
				},
			},
		}
		for i, test := range tests {
			e := headerParamEncoder{
				receiver:  newReceiver(),
				paramName: test.HeaderName,
				explode:   test.Explode,
				header:    make(http.Header),
			}
			require.NoError(t, e.EncodeArray(func(e Encoder) error {
				for _, v := range test.Input {
					if err := e.EncodeValue(v); err != nil {
						return err
					}
				}
				return nil
			}))
			require.NoError(t, e.serialize())
			require.Equal(t, test.Expect, e.header, fmt.Sprintf("Test %d", i+1))
		}
	})

	t.Run("object", func(t *testing.T) {
		tests := []struct {
			HeaderName string
			Input      []Field
			Explode    bool
			Expect     http.Header
		}{
			{
				HeaderName: "X-MyHeader",
				Input: []Field{
					{
						Name:  "role",
						Value: "admin",
					},
					{
						Name:  "firstName",
						Value: "Alex",
					},
				},
				Explode: false,
				Expect: http.Header{
					"X-Myheader": []string{"role,admin,firstName,Alex"},
				},
			},
			{
				HeaderName: "X-MyHeader",
				Input: []Field{
					{
						Name:  "role",
						Value: "admin",
					},
					{
						Name:  "firstName",
						Value: "Alex",
					},
				},
				Explode: true,
				Expect: http.Header{
					"X-Myheader": []string{"role=admin,firstName=Alex"},
				},
			},
		}
		for i, test := range tests {
			e := headerParamEncoder{
				receiver:  newReceiver(),
				paramName: test.HeaderName,
				explode:   test.Explode,
				header:    make(http.Header),
			}
			for _, f := range test.Input {
				require.NoError(t, e.EncodeField(f.Name, func(e Encoder) error {
					return e.EncodeValue(f.Value)
				}))
			}
			require.NoError(t, e.serialize())
			require.Equal(t, test.Expect, e.header, fmt.Sprintf("Test %d", i+1))
		}
	})
}
