// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
)

func (s *Server) notFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	if len(elem) == 0 {
		s.notFound(w, r)
		return
	}

	args := map[string]string{}
	// Static code generated router with unwrapped path search.
	switch r.Method {
	case "POST":
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				s.handleBanChatMemberRequest(args, w, r)
				return
			}
			switch elem[0] {
			case 'a': // Prefix: "a"
				if l := len("a"); len(elem) >= l && elem[0:l] == "a" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleAnswerCallbackQueryRequest(args, w, r)
					return
				}
				switch elem[0] {
				case 'd': // Prefix: "ddStickerToSet"
					if l := len("ddStickerToSet"); len(elem) >= l && elem[0:l] == "ddStickerToSet" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: AddStickerToSet
						s.handleAddStickerToSetRequest(args, w, r)
						return
					}
				case 'n': // Prefix: "nswer"
					if l := len("nswer"); len(elem) >= l && elem[0:l] == "nswer" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						s.handleAnswerInlineQueryRequest(args, w, r)
						return
					}
					switch elem[0] {
					case 'C': // Prefix: "CallbackQuery"
						if l := len("CallbackQuery"); len(elem) >= l && elem[0:l] == "CallbackQuery" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: AnswerCallbackQuery
							s.handleAnswerCallbackQueryRequest(args, w, r)
							return
						}
					case 'I': // Prefix: "InlineQuery"
						if l := len("InlineQuery"); len(elem) >= l && elem[0:l] == "InlineQuery" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: AnswerInlineQuery
							s.handleAnswerInlineQueryRequest(args, w, r)
							return
						}
					case 'P': // Prefix: "PreCheckoutQuery"
						if l := len("PreCheckoutQuery"); len(elem) >= l && elem[0:l] == "PreCheckoutQuery" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: AnswerPreCheckoutQuery
							s.handleAnswerPreCheckoutQueryRequest(args, w, r)
							return
						}
					case 'S': // Prefix: "ShippingQuery"
						if l := len("ShippingQuery"); len(elem) >= l && elem[0:l] == "ShippingQuery" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: AnswerShippingQuery
							s.handleAnswerShippingQueryRequest(args, w, r)
							return
						}
					}
				case 'p': // Prefix: "pproveChatJoinRequest"
					if l := len("pproveChatJoinRequest"); len(elem) >= l && elem[0:l] == "pproveChatJoinRequest" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: ApproveChatJoinRequest
						s.handleApproveChatJoinRequestRequest(args, w, r)
						return
					}
				}
			case 'b': // Prefix: "banChatMember"
				if l := len("banChatMember"); len(elem) >= l && elem[0:l] == "banChatMember" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: BanChatMember
					s.handleBanChatMemberRequest(args, w, r)
					return
				}
			case 'c': // Prefix: "c"
				if l := len("c"); len(elem) >= l && elem[0:l] == "c" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleCreateChatInviteLinkRequest(args, w, r)
					return
				}
				switch elem[0] {
				case 'o': // Prefix: "opyMessage"
					if l := len("opyMessage"); len(elem) >= l && elem[0:l] == "opyMessage" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: CopyMessage
						s.handleCopyMessageRequest(args, w, r)
						return
					}
				case 'r': // Prefix: "reate"
					if l := len("reate"); len(elem) >= l && elem[0:l] == "reate" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						s.handleCreateNewStickerSetRequest(args, w, r)
						return
					}
					switch elem[0] {
					case 'C': // Prefix: "ChatInviteLink"
						if l := len("ChatInviteLink"); len(elem) >= l && elem[0:l] == "ChatInviteLink" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: CreateChatInviteLink
							s.handleCreateChatInviteLinkRequest(args, w, r)
							return
						}
					case 'N': // Prefix: "NewStickerSet"
						if l := len("NewStickerSet"); len(elem) >= l && elem[0:l] == "NewStickerSet" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: CreateNewStickerSet
							s.handleCreateNewStickerSetRequest(args, w, r)
							return
						}
					}
				}
			case 'd': // Prefix: "de"
				if l := len("de"); len(elem) >= l && elem[0:l] == "de" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleDeleteChatPhotoRequest(args, w, r)
					return
				}
				switch elem[0] {
				case 'c': // Prefix: "clineChatJoinRequest"
					if l := len("clineChatJoinRequest"); len(elem) >= l && elem[0:l] == "clineChatJoinRequest" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: DeclineChatJoinRequest
						s.handleDeclineChatJoinRequestRequest(args, w, r)
						return
					}
				case 'l': // Prefix: "lete"
					if l := len("lete"); len(elem) >= l && elem[0:l] == "lete" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						s.handleDeleteMessageRequest(args, w, r)
						return
					}
					switch elem[0] {
					case 'C': // Prefix: "Chat"
						if l := len("Chat"); len(elem) >= l && elem[0:l] == "Chat" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							s.handleDeleteChatStickerSetRequest(args, w, r)
							return
						}
						switch elem[0] {
						case 'P': // Prefix: "Photo"
							if l := len("Photo"); len(elem) >= l && elem[0:l] == "Photo" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: DeleteChatPhoto
								s.handleDeleteChatPhotoRequest(args, w, r)
								return
							}
						case 'S': // Prefix: "StickerSet"
							if l := len("StickerSet"); len(elem) >= l && elem[0:l] == "StickerSet" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: DeleteChatStickerSet
								s.handleDeleteChatStickerSetRequest(args, w, r)
								return
							}
						}
					case 'M': // Prefix: "M"
						if l := len("M"); len(elem) >= l && elem[0:l] == "M" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							s.handleDeleteMyCommandsRequest(args, w, r)
							return
						}
						switch elem[0] {
						case 'e': // Prefix: "essage"
							if l := len("essage"); len(elem) >= l && elem[0:l] == "essage" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: DeleteMessage
								s.handleDeleteMessageRequest(args, w, r)
								return
							}
						case 'y': // Prefix: "yCommands"
							if l := len("yCommands"); len(elem) >= l && elem[0:l] == "yCommands" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: DeleteMyCommands
								s.handleDeleteMyCommandsRequest(args, w, r)
								return
							}
						}
					case 'S': // Prefix: "StickerFromSet"
						if l := len("StickerFromSet"); len(elem) >= l && elem[0:l] == "StickerFromSet" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: DeleteStickerFromSet
							s.handleDeleteStickerFromSetRequest(args, w, r)
							return
						}
					case 'W': // Prefix: "Webhook"
						if l := len("Webhook"); len(elem) >= l && elem[0:l] == "Webhook" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: DeleteWebhook
							s.handleDeleteWebhookRequest(args, w, r)
							return
						}
					}
				}
			case 'e': // Prefix: "e"
				if l := len("e"); len(elem) >= l && elem[0:l] == "e" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleExportChatInviteLinkRequest(args, w, r)
					return
				}
				switch elem[0] {
				case 'd': // Prefix: "dit"
					if l := len("dit"); len(elem) >= l && elem[0:l] == "dit" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						s.handleEditMessageCaptionRequest(args, w, r)
						return
					}
					switch elem[0] {
					case 'C': // Prefix: "ChatInviteLink"
						if l := len("ChatInviteLink"); len(elem) >= l && elem[0:l] == "ChatInviteLink" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: EditChatInviteLink
							s.handleEditChatInviteLinkRequest(args, w, r)
							return
						}
					case 'M': // Prefix: "Message"
						if l := len("Message"); len(elem) >= l && elem[0:l] == "Message" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							s.handleEditMessageLiveLocationRequest(args, w, r)
							return
						}
						switch elem[0] {
						case 'C': // Prefix: "Caption"
							if l := len("Caption"); len(elem) >= l && elem[0:l] == "Caption" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: EditMessageCaption
								s.handleEditMessageCaptionRequest(args, w, r)
								return
							}
						case 'L': // Prefix: "LiveLocation"
							if l := len("LiveLocation"); len(elem) >= l && elem[0:l] == "LiveLocation" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: EditMessageLiveLocation
								s.handleEditMessageLiveLocationRequest(args, w, r)
								return
							}
						case 'M': // Prefix: "Media"
							if l := len("Media"); len(elem) >= l && elem[0:l] == "Media" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: EditMessageMedia
								s.handleEditMessageMediaRequest(args, w, r)
								return
							}
						case 'R': // Prefix: "ReplyMarkup"
							if l := len("ReplyMarkup"); len(elem) >= l && elem[0:l] == "ReplyMarkup" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: EditMessageReplyMarkup
								s.handleEditMessageReplyMarkupRequest(args, w, r)
								return
							}
						case 'T': // Prefix: "Text"
							if l := len("Text"); len(elem) >= l && elem[0:l] == "Text" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: EditMessageText
								s.handleEditMessageTextRequest(args, w, r)
								return
							}
						}
					}
				case 'x': // Prefix: "xportChatInviteLink"
					if l := len("xportChatInviteLink"); len(elem) >= l && elem[0:l] == "xportChatInviteLink" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: ExportChatInviteLink
						s.handleExportChatInviteLinkRequest(args, w, r)
						return
					}
				}
			case 'f': // Prefix: "forwardMessage"
				if l := len("forwardMessage"); len(elem) >= l && elem[0:l] == "forwardMessage" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: ForwardMessage
					s.handleForwardMessageRequest(args, w, r)
					return
				}
			case 'g': // Prefix: "get"
				if l := len("get"); len(elem) >= l && elem[0:l] == "get" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleGetFileRequest(args, w, r)
					return
				}
				switch elem[0] {
				case 'C': // Prefix: "Chat"
					if l := len("Chat"); len(elem) >= l && elem[0:l] == "Chat" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						s.handleGetChatRequest(args, w, r)
						return
					}
					switch elem[0] {
					case 'A': // Prefix: "Administrators"
						if l := len("Administrators"); len(elem) >= l && elem[0:l] == "Administrators" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: GetChatAdministrators
							s.handleGetChatAdministratorsRequest(args, w, r)
							return
						}
					case 'M': // Prefix: "Member"
						if l := len("Member"); len(elem) >= l && elem[0:l] == "Member" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							s.handleGetChatMemberRequest(args, w, r)
							return
						}
						switch elem[0] {
						case 'C': // Prefix: "Count"
							if l := len("Count"); len(elem) >= l && elem[0:l] == "Count" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: GetChatMemberCount
								s.handleGetChatMemberCountRequest(args, w, r)
								return
							}
						}
					}
				case 'F': // Prefix: "File"
					if l := len("File"); len(elem) >= l && elem[0:l] == "File" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: GetFile
						s.handleGetFileRequest(args, w, r)
						return
					}
				case 'G': // Prefix: "GameHighScores"
					if l := len("GameHighScores"); len(elem) >= l && elem[0:l] == "GameHighScores" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: GetGameHighScores
						s.handleGetGameHighScoresRequest(args, w, r)
						return
					}
				case 'M': // Prefix: "M"
					if l := len("M"); len(elem) >= l && elem[0:l] == "M" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						s.handleGetMyCommandsRequest(args, w, r)
						return
					}
					switch elem[0] {
					case 'e': // Prefix: "e"
						if l := len("e"); len(elem) >= l && elem[0:l] == "e" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: GetMe
							s.handleGetMeRequest(args, w, r)
							return
						}
					case 'y': // Prefix: "yCommands"
						if l := len("yCommands"); len(elem) >= l && elem[0:l] == "yCommands" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: GetMyCommands
							s.handleGetMyCommandsRequest(args, w, r)
							return
						}
					}
				case 'S': // Prefix: "StickerSet"
					if l := len("StickerSet"); len(elem) >= l && elem[0:l] == "StickerSet" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: GetStickerSet
						s.handleGetStickerSetRequest(args, w, r)
						return
					}
				case 'U': // Prefix: "U"
					if l := len("U"); len(elem) >= l && elem[0:l] == "U" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						s.handleGetUserProfilePhotosRequest(args, w, r)
						return
					}
					switch elem[0] {
					case 'p': // Prefix: "pdates"
						if l := len("pdates"); len(elem) >= l && elem[0:l] == "pdates" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: GetUpdates
							s.handleGetUpdatesRequest(args, w, r)
							return
						}
					case 's': // Prefix: "serProfilePhotos"
						if l := len("serProfilePhotos"); len(elem) >= l && elem[0:l] == "serProfilePhotos" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: GetUserProfilePhotos
							s.handleGetUserProfilePhotosRequest(args, w, r)
							return
						}
					}
				}
			case 'l': // Prefix: "leaveChat"
				if l := len("leaveChat"); len(elem) >= l && elem[0:l] == "leaveChat" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf: LeaveChat
					s.handleLeaveChatRequest(args, w, r)
					return
				}
			case 'p': // Prefix: "p"
				if l := len("p"); len(elem) >= l && elem[0:l] == "p" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handlePromoteChatMemberRequest(args, w, r)
					return
				}
				switch elem[0] {
				case 'i': // Prefix: "inChatMessage"
					if l := len("inChatMessage"); len(elem) >= l && elem[0:l] == "inChatMessage" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: PinChatMessage
						s.handlePinChatMessageRequest(args, w, r)
						return
					}
				case 'r': // Prefix: "romoteChatMember"
					if l := len("romoteChatMember"); len(elem) >= l && elem[0:l] == "romoteChatMember" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: PromoteChatMember
						s.handlePromoteChatMemberRequest(args, w, r)
						return
					}
				}
			case 'r': // Prefix: "re"
				if l := len("re"); len(elem) >= l && elem[0:l] == "re" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleRevokeChatInviteLinkRequest(args, w, r)
					return
				}
				switch elem[0] {
				case 's': // Prefix: "strictChatMember"
					if l := len("strictChatMember"); len(elem) >= l && elem[0:l] == "strictChatMember" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: RestrictChatMember
						s.handleRestrictChatMemberRequest(args, w, r)
						return
					}
				case 'v': // Prefix: "vokeChatInviteLink"
					if l := len("vokeChatInviteLink"); len(elem) >= l && elem[0:l] == "vokeChatInviteLink" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: RevokeChatInviteLink
						s.handleRevokeChatInviteLinkRequest(args, w, r)
						return
					}
				}
			case 's': // Prefix: "s"
				if l := len("s"); len(elem) >= l && elem[0:l] == "s" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleStopMessageLiveLocationRequest(args, w, r)
					return
				}
				switch elem[0] {
				case 'e': // Prefix: "e"
					if l := len("e"); len(elem) >= l && elem[0:l] == "e" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						s.handleSetChatAdministratorCustomTitleRequest(args, w, r)
						return
					}
					switch elem[0] {
					case 'n': // Prefix: "nd"
						if l := len("nd"); len(elem) >= l && elem[0:l] == "nd" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							s.handleSendChatActionRequest(args, w, r)
							return
						}
						switch elem[0] {
						case 'A': // Prefix: "A"
							if l := len("A"); len(elem) >= l && elem[0:l] == "A" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								s.handleSendAudioRequest(args, w, r)
								return
							}
							switch elem[0] {
							case 'n': // Prefix: "nimation"
								if l := len("nimation"); len(elem) >= l && elem[0:l] == "nimation" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf: SendAnimation
									s.handleSendAnimationRequest(args, w, r)
									return
								}
							case 'u': // Prefix: "udio"
								if l := len("udio"); len(elem) >= l && elem[0:l] == "udio" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf: SendAudio
									s.handleSendAudioRequest(args, w, r)
									return
								}
							}
						case 'C': // Prefix: "C"
							if l := len("C"); len(elem) >= l && elem[0:l] == "C" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								s.handleSendContactRequest(args, w, r)
								return
							}
							switch elem[0] {
							case 'h': // Prefix: "hatAction"
								if l := len("hatAction"); len(elem) >= l && elem[0:l] == "hatAction" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf: SendChatAction
									s.handleSendChatActionRequest(args, w, r)
									return
								}
							case 'o': // Prefix: "ontact"
								if l := len("ontact"); len(elem) >= l && elem[0:l] == "ontact" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf: SendContact
									s.handleSendContactRequest(args, w, r)
									return
								}
							}
						case 'D': // Prefix: "D"
							if l := len("D"); len(elem) >= l && elem[0:l] == "D" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								s.handleSendDocumentRequest(args, w, r)
								return
							}
							switch elem[0] {
							case 'i': // Prefix: "ice"
								if l := len("ice"); len(elem) >= l && elem[0:l] == "ice" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf: SendDice
									s.handleSendDiceRequest(args, w, r)
									return
								}
							case 'o': // Prefix: "ocument"
								if l := len("ocument"); len(elem) >= l && elem[0:l] == "ocument" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf: SendDocument
									s.handleSendDocumentRequest(args, w, r)
									return
								}
							}
						case 'G': // Prefix: "Game"
							if l := len("Game"); len(elem) >= l && elem[0:l] == "Game" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: SendGame
								s.handleSendGameRequest(args, w, r)
								return
							}
						case 'I': // Prefix: "Invoice"
							if l := len("Invoice"); len(elem) >= l && elem[0:l] == "Invoice" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: SendInvoice
								s.handleSendInvoiceRequest(args, w, r)
								return
							}
						case 'L': // Prefix: "Location"
							if l := len("Location"); len(elem) >= l && elem[0:l] == "Location" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: SendLocation
								s.handleSendLocationRequest(args, w, r)
								return
							}
						case 'M': // Prefix: "Me"
							if l := len("Me"); len(elem) >= l && elem[0:l] == "Me" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								s.handleSendMessageRequest(args, w, r)
								return
							}
							switch elem[0] {
							case 'd': // Prefix: "diaGroup"
								if l := len("diaGroup"); len(elem) >= l && elem[0:l] == "diaGroup" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf: SendMediaGroup
									s.handleSendMediaGroupRequest(args, w, r)
									return
								}
							case 's': // Prefix: "ssage"
								if l := len("ssage"); len(elem) >= l && elem[0:l] == "ssage" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf: SendMessage
									s.handleSendMessageRequest(args, w, r)
									return
								}
							}
						case 'P': // Prefix: "P"
							if l := len("P"); len(elem) >= l && elem[0:l] == "P" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								s.handleSendPollRequest(args, w, r)
								return
							}
							switch elem[0] {
							case 'h': // Prefix: "hoto"
								if l := len("hoto"); len(elem) >= l && elem[0:l] == "hoto" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf: SendPhoto
									s.handleSendPhotoRequest(args, w, r)
									return
								}
							case 'o': // Prefix: "oll"
								if l := len("oll"); len(elem) >= l && elem[0:l] == "oll" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf: SendPoll
									s.handleSendPollRequest(args, w, r)
									return
								}
							}
						case 'S': // Prefix: "Sticker"
							if l := len("Sticker"); len(elem) >= l && elem[0:l] == "Sticker" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: SendSticker
								s.handleSendStickerRequest(args, w, r)
								return
							}
						case 'V': // Prefix: "V"
							if l := len("V"); len(elem) >= l && elem[0:l] == "V" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								s.handleSendVideoRequest(args, w, r)
								return
							}
							switch elem[0] {
							case 'e': // Prefix: "enue"
								if l := len("enue"); len(elem) >= l && elem[0:l] == "enue" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf: SendVenue
									s.handleSendVenueRequest(args, w, r)
									return
								}
							case 'i': // Prefix: "ideo"
								if l := len("ideo"); len(elem) >= l && elem[0:l] == "ideo" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									s.handleSendVideoRequest(args, w, r)
									return
								}
								switch elem[0] {
								case 'N': // Prefix: "Note"
									if l := len("Note"); len(elem) >= l && elem[0:l] == "Note" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf: SendVideoNote
										s.handleSendVideoNoteRequest(args, w, r)
										return
									}
								}
							case 'o': // Prefix: "oice"
								if l := len("oice"); len(elem) >= l && elem[0:l] == "oice" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf: SendVoice
									s.handleSendVoiceRequest(args, w, r)
									return
								}
							}
						}
					case 't': // Prefix: "t"
						if l := len("t"); len(elem) >= l && elem[0:l] == "t" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							s.handleSetGameScoreRequest(args, w, r)
							return
						}
						switch elem[0] {
						case 'C': // Prefix: "Chat"
							if l := len("Chat"); len(elem) >= l && elem[0:l] == "Chat" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								s.handleSetChatDescriptionRequest(args, w, r)
								return
							}
							switch elem[0] {
							case 'A': // Prefix: "AdministratorCustomTitle"
								if l := len("AdministratorCustomTitle"); len(elem) >= l && elem[0:l] == "AdministratorCustomTitle" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf: SetChatAdministratorCustomTitle
									s.handleSetChatAdministratorCustomTitleRequest(args, w, r)
									return
								}
							case 'D': // Prefix: "Description"
								if l := len("Description"); len(elem) >= l && elem[0:l] == "Description" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf: SetChatDescription
									s.handleSetChatDescriptionRequest(args, w, r)
									return
								}
							case 'P': // Prefix: "P"
								if l := len("P"); len(elem) >= l && elem[0:l] == "P" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									s.handleSetChatPhotoRequest(args, w, r)
									return
								}
								switch elem[0] {
								case 'e': // Prefix: "ermissions"
									if l := len("ermissions"); len(elem) >= l && elem[0:l] == "ermissions" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf: SetChatPermissions
										s.handleSetChatPermissionsRequest(args, w, r)
										return
									}
								case 'h': // Prefix: "hoto"
									if l := len("hoto"); len(elem) >= l && elem[0:l] == "hoto" {
										elem = elem[l:]
									} else {
										break
									}

									if len(elem) == 0 {
										// Leaf: SetChatPhoto
										s.handleSetChatPhotoRequest(args, w, r)
										return
									}
								}
							case 'S': // Prefix: "StickerSet"
								if l := len("StickerSet"); len(elem) >= l && elem[0:l] == "StickerSet" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf: SetChatStickerSet
									s.handleSetChatStickerSetRequest(args, w, r)
									return
								}
							case 'T': // Prefix: "Title"
								if l := len("Title"); len(elem) >= l && elem[0:l] == "Title" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf: SetChatTitle
									s.handleSetChatTitleRequest(args, w, r)
									return
								}
							}
						case 'G': // Prefix: "GameScore"
							if l := len("GameScore"); len(elem) >= l && elem[0:l] == "GameScore" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: SetGameScore
								s.handleSetGameScoreRequest(args, w, r)
								return
							}
						case 'M': // Prefix: "MyCommands"
							if l := len("MyCommands"); len(elem) >= l && elem[0:l] == "MyCommands" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: SetMyCommands
								s.handleSetMyCommandsRequest(args, w, r)
								return
							}
						case 'P': // Prefix: "PassportDataErrors"
							if l := len("PassportDataErrors"); len(elem) >= l && elem[0:l] == "PassportDataErrors" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: SetPassportDataErrors
								s.handleSetPassportDataErrorsRequest(args, w, r)
								return
							}
						case 'S': // Prefix: "Sticker"
							if l := len("Sticker"); len(elem) >= l && elem[0:l] == "Sticker" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								s.handleSetStickerSetThumbRequest(args, w, r)
								return
							}
							switch elem[0] {
							case 'P': // Prefix: "PositionInSet"
								if l := len("PositionInSet"); len(elem) >= l && elem[0:l] == "PositionInSet" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf: SetStickerPositionInSet
									s.handleSetStickerPositionInSetRequest(args, w, r)
									return
								}
							case 'S': // Prefix: "SetThumb"
								if l := len("SetThumb"); len(elem) >= l && elem[0:l] == "SetThumb" {
									elem = elem[l:]
								} else {
									break
								}

								if len(elem) == 0 {
									// Leaf: SetStickerSetThumb
									s.handleSetStickerSetThumbRequest(args, w, r)
									return
								}
							}
						case 'W': // Prefix: "Webhook"
							if l := len("Webhook"); len(elem) >= l && elem[0:l] == "Webhook" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: SetWebhook
								s.handleSetWebhookRequest(args, w, r)
								return
							}
						}
					}
				case 't': // Prefix: "top"
					if l := len("top"); len(elem) >= l && elem[0:l] == "top" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						s.handleStopPollRequest(args, w, r)
						return
					}
					switch elem[0] {
					case 'M': // Prefix: "MessageLiveLocation"
						if l := len("MessageLiveLocation"); len(elem) >= l && elem[0:l] == "MessageLiveLocation" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: StopMessageLiveLocation
							s.handleStopMessageLiveLocationRequest(args, w, r)
							return
						}
					case 'P': // Prefix: "Poll"
						if l := len("Poll"); len(elem) >= l && elem[0:l] == "Poll" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: StopPoll
							s.handleStopPollRequest(args, w, r)
							return
						}
					}
				}
			case 'u': // Prefix: "u"
				if l := len("u"); len(elem) >= l && elem[0:l] == "u" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					s.handleUploadStickerFileRequest(args, w, r)
					return
				}
				switch elem[0] {
				case 'n': // Prefix: "n"
					if l := len("n"); len(elem) >= l && elem[0:l] == "n" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						s.handleUnpinAllChatMessagesRequest(args, w, r)
						return
					}
					switch elem[0] {
					case 'b': // Prefix: "banChatMember"
						if l := len("banChatMember"); len(elem) >= l && elem[0:l] == "banChatMember" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf: UnbanChatMember
							s.handleUnbanChatMemberRequest(args, w, r)
							return
						}
					case 'p': // Prefix: "pin"
						if l := len("pin"); len(elem) >= l && elem[0:l] == "pin" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							s.handleUnpinChatMessageRequest(args, w, r)
							return
						}
						switch elem[0] {
						case 'A': // Prefix: "AllChatMessages"
							if l := len("AllChatMessages"); len(elem) >= l && elem[0:l] == "AllChatMessages" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: UnpinAllChatMessages
								s.handleUnpinAllChatMessagesRequest(args, w, r)
								return
							}
						case 'C': // Prefix: "ChatMessage"
							if l := len("ChatMessage"); len(elem) >= l && elem[0:l] == "ChatMessage" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf: UnpinChatMessage
								s.handleUnpinChatMessageRequest(args, w, r)
								return
							}
						}
					}
				case 'p': // Prefix: "ploadStickerFile"
					if l := len("ploadStickerFile"); len(elem) >= l && elem[0:l] == "ploadStickerFile" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf: UploadStickerFile
						s.handleUploadStickerFileRequest(args, w, r)
						return
					}
				}
			}
		}
	}
	s.notFound(w, r)
}
