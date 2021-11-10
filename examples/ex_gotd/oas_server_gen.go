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

	"github.com/go-chi/chi/v5"
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
	_ = chi.Context{}
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

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// AddStickerToSet implements addStickerToSet operation.
	//
	// POST /addStickerToSet
	AddStickerToSet(ctx context.Context, req AddStickerToSet) (AddStickerToSetRes, error)
	// AnswerCallbackQuery implements answerCallbackQuery operation.
	//
	// POST /answerCallbackQuery
	AnswerCallbackQuery(ctx context.Context, req AnswerCallbackQuery) (AnswerCallbackQueryRes, error)
	// AnswerInlineQuery implements answerInlineQuery operation.
	//
	// POST /answerInlineQuery
	AnswerInlineQuery(ctx context.Context, req AnswerInlineQuery) (AnswerInlineQueryRes, error)
	// AnswerPreCheckoutQuery implements answerPreCheckoutQuery operation.
	//
	// POST /answerPreCheckoutQuery
	AnswerPreCheckoutQuery(ctx context.Context, req AnswerPreCheckoutQuery) (AnswerPreCheckoutQueryRes, error)
	// AnswerShippingQuery implements answerShippingQuery operation.
	//
	// POST /answerShippingQuery
	AnswerShippingQuery(ctx context.Context, req AnswerShippingQuery) (AnswerShippingQueryRes, error)
	// BanChatMember implements banChatMember operation.
	//
	// POST /banChatMember
	BanChatMember(ctx context.Context, req BanChatMember) (BanChatMemberRes, error)
	// CopyMessage implements copyMessage operation.
	//
	// POST /copyMessage
	CopyMessage(ctx context.Context, req CopyMessage) (CopyMessageRes, error)
	// CreateChatInviteLink implements createChatInviteLink operation.
	//
	// POST /createChatInviteLink
	CreateChatInviteLink(ctx context.Context, req CreateChatInviteLink) (CreateChatInviteLinkRes, error)
	// CreateNewStickerSet implements createNewStickerSet operation.
	//
	// POST /createNewStickerSet
	CreateNewStickerSet(ctx context.Context, req CreateNewStickerSet) (CreateNewStickerSetRes, error)
	// DeleteChatPhoto implements deleteChatPhoto operation.
	//
	// POST /deleteChatPhoto
	DeleteChatPhoto(ctx context.Context, req DeleteChatPhoto) (DeleteChatPhotoRes, error)
	// DeleteChatStickerSet implements deleteChatStickerSet operation.
	//
	// POST /deleteChatStickerSet
	DeleteChatStickerSet(ctx context.Context, req DeleteChatStickerSet) (DeleteChatStickerSetRes, error)
	// DeleteMessage implements deleteMessage operation.
	//
	// POST /deleteMessage
	DeleteMessage(ctx context.Context, req DeleteMessage) (DeleteMessageRes, error)
	// DeleteMyCommands implements deleteMyCommands operation.
	//
	// POST /deleteMyCommands
	DeleteMyCommands(ctx context.Context, req DeleteMyCommands) (DeleteMyCommandsRes, error)
	// DeleteStickerFromSet implements deleteStickerFromSet operation.
	//
	// POST /deleteStickerFromSet
	DeleteStickerFromSet(ctx context.Context, req DeleteStickerFromSet) (DeleteStickerFromSetRes, error)
	// DeleteWebhook implements deleteWebhook operation.
	//
	// POST /deleteWebhook
	DeleteWebhook(ctx context.Context, req DeleteWebhook) (DeleteWebhookRes, error)
	// EditChatInviteLink implements editChatInviteLink operation.
	//
	// POST /editChatInviteLink
	EditChatInviteLink(ctx context.Context, req EditChatInviteLink) (EditChatInviteLinkRes, error)
	// EditMessageCaption implements editMessageCaption operation.
	//
	// POST /editMessageCaption
	EditMessageCaption(ctx context.Context, req EditMessageCaption) (EditMessageCaptionRes, error)
	// EditMessageLiveLocation implements editMessageLiveLocation operation.
	//
	// POST /editMessageLiveLocation
	EditMessageLiveLocation(ctx context.Context, req EditMessageLiveLocation) (EditMessageLiveLocationRes, error)
	// EditMessageMedia implements editMessageMedia operation.
	//
	// POST /editMessageMedia
	EditMessageMedia(ctx context.Context, req EditMessageMedia) (EditMessageMediaRes, error)
	// EditMessageReplyMarkup implements editMessageReplyMarkup operation.
	//
	// POST /editMessageReplyMarkup
	EditMessageReplyMarkup(ctx context.Context, req EditMessageReplyMarkup) (EditMessageReplyMarkupRes, error)
	// EditMessageText implements editMessageText operation.
	//
	// POST /editMessageText
	EditMessageText(ctx context.Context, req EditMessageText) (EditMessageTextRes, error)
	// ExportChatInviteLink implements exportChatInviteLink operation.
	//
	// POST /exportChatInviteLink
	ExportChatInviteLink(ctx context.Context, req ExportChatInviteLink) (ExportChatInviteLinkRes, error)
	// ForwardMessage implements forwardMessage operation.
	//
	// POST /forwardMessage
	ForwardMessage(ctx context.Context, req ForwardMessage) (ForwardMessageRes, error)
	// GetChat implements getChat operation.
	//
	// POST /getChat
	GetChat(ctx context.Context, req GetChat) (GetChatRes, error)
	// GetChatAdministrators implements getChatAdministrators operation.
	//
	// POST /getChatAdministrators
	GetChatAdministrators(ctx context.Context, req GetChatAdministrators) (GetChatAdministratorsRes, error)
	// GetChatMember implements getChatMember operation.
	//
	// POST /getChatMember
	GetChatMember(ctx context.Context, req GetChatMember) (GetChatMemberRes, error)
	// GetChatMemberCount implements getChatMemberCount operation.
	//
	// POST /getChatMemberCount
	GetChatMemberCount(ctx context.Context, req GetChatMemberCount) (GetChatMemberCountRes, error)
	// GetFile implements getFile operation.
	//
	// POST /getFile
	GetFile(ctx context.Context, req GetFile) (GetFileRes, error)
	// GetGameHighScores implements getGameHighScores operation.
	//
	// POST /getGameHighScores
	GetGameHighScores(ctx context.Context, req GetGameHighScores) (GetGameHighScoresRes, error)
	// GetMe implements getMe operation.
	//
	// POST /getMe
	GetMe(ctx context.Context) (GetMeRes, error)
	// GetMyCommands implements getMyCommands operation.
	//
	// POST /getMyCommands
	GetMyCommands(ctx context.Context, req GetMyCommands) (GetMyCommandsRes, error)
	// GetStickerSet implements getStickerSet operation.
	//
	// POST /getStickerSet
	GetStickerSet(ctx context.Context, req GetStickerSet) (GetStickerSetRes, error)
	// GetUpdates implements getUpdates operation.
	//
	// POST /getUpdates
	GetUpdates(ctx context.Context, req GetUpdates) (GetUpdatesRes, error)
	// GetUserProfilePhotos implements getUserProfilePhotos operation.
	//
	// POST /getUserProfilePhotos
	GetUserProfilePhotos(ctx context.Context, req GetUserProfilePhotos) (GetUserProfilePhotosRes, error)
	// LeaveChat implements leaveChat operation.
	//
	// POST /leaveChat
	LeaveChat(ctx context.Context, req LeaveChat) (LeaveChatRes, error)
	// PinChatMessage implements pinChatMessage operation.
	//
	// POST /pinChatMessage
	PinChatMessage(ctx context.Context, req PinChatMessage) (PinChatMessageRes, error)
	// PromoteChatMember implements promoteChatMember operation.
	//
	// POST /promoteChatMember
	PromoteChatMember(ctx context.Context, req PromoteChatMember) (PromoteChatMemberRes, error)
	// RestrictChatMember implements restrictChatMember operation.
	//
	// POST /restrictChatMember
	RestrictChatMember(ctx context.Context, req RestrictChatMember) (RestrictChatMemberRes, error)
	// RevokeChatInviteLink implements revokeChatInviteLink operation.
	//
	// POST /revokeChatInviteLink
	RevokeChatInviteLink(ctx context.Context, req RevokeChatInviteLink) (RevokeChatInviteLinkRes, error)
	// SendAnimation implements sendAnimation operation.
	//
	// POST /sendAnimation
	SendAnimation(ctx context.Context, req SendAnimation) (SendAnimationRes, error)
	// SendAudio implements sendAudio operation.
	//
	// POST /sendAudio
	SendAudio(ctx context.Context, req SendAudio) (SendAudioRes, error)
	// SendChatAction implements sendChatAction operation.
	//
	// POST /sendChatAction
	SendChatAction(ctx context.Context, req SendChatAction) (SendChatActionRes, error)
	// SendContact implements sendContact operation.
	//
	// POST /sendContact
	SendContact(ctx context.Context, req SendContact) (SendContactRes, error)
	// SendDice implements sendDice operation.
	//
	// POST /sendDice
	SendDice(ctx context.Context, req SendDice) (SendDiceRes, error)
	// SendDocument implements sendDocument operation.
	//
	// POST /sendDocument
	SendDocument(ctx context.Context, req SendDocument) (SendDocumentRes, error)
	// SendGame implements sendGame operation.
	//
	// POST /sendGame
	SendGame(ctx context.Context, req SendGame) (SendGameRes, error)
	// SendInvoice implements sendInvoice operation.
	//
	// POST /sendInvoice
	SendInvoice(ctx context.Context, req SendInvoice) (SendInvoiceRes, error)
	// SendLocation implements sendLocation operation.
	//
	// POST /sendLocation
	SendLocation(ctx context.Context, req SendLocation) (SendLocationRes, error)
	// SendMediaGroup implements sendMediaGroup operation.
	//
	// POST /sendMediaGroup
	SendMediaGroup(ctx context.Context, req SendMediaGroup) (SendMediaGroupRes, error)
	// SendMessage implements sendMessage operation.
	//
	// POST /sendMessage
	SendMessage(ctx context.Context, req SendMessage) (SendMessageRes, error)
	// SendPhoto implements sendPhoto operation.
	//
	// POST /sendPhoto
	SendPhoto(ctx context.Context, req SendPhoto) (SendPhotoRes, error)
	// SendPoll implements sendPoll operation.
	//
	// POST /sendPoll
	SendPoll(ctx context.Context, req SendPoll) (SendPollRes, error)
	// SendSticker implements sendSticker operation.
	//
	// POST /sendSticker
	SendSticker(ctx context.Context, req SendSticker) (SendStickerRes, error)
	// SendVenue implements sendVenue operation.
	//
	// POST /sendVenue
	SendVenue(ctx context.Context, req SendVenue) (SendVenueRes, error)
	// SendVideo implements sendVideo operation.
	//
	// POST /sendVideo
	SendVideo(ctx context.Context, req SendVideo) (SendVideoRes, error)
	// SendVideoNote implements sendVideoNote operation.
	//
	// POST /sendVideoNote
	SendVideoNote(ctx context.Context, req SendVideoNote) (SendVideoNoteRes, error)
	// SendVoice implements sendVoice operation.
	//
	// POST /sendVoice
	SendVoice(ctx context.Context, req SendVoice) (SendVoiceRes, error)
	// SetChatAdministratorCustomTitle implements setChatAdministratorCustomTitle operation.
	//
	// POST /setChatAdministratorCustomTitle
	SetChatAdministratorCustomTitle(ctx context.Context, req SetChatAdministratorCustomTitle) (SetChatAdministratorCustomTitleRes, error)
	// SetChatDescription implements setChatDescription operation.
	//
	// POST /setChatDescription
	SetChatDescription(ctx context.Context, req SetChatDescription) (SetChatDescriptionRes, error)
	// SetChatPermissions implements setChatPermissions operation.
	//
	// POST /setChatPermissions
	SetChatPermissions(ctx context.Context, req SetChatPermissions) (SetChatPermissionsRes, error)
	// SetChatPhoto implements setChatPhoto operation.
	//
	// POST /setChatPhoto
	SetChatPhoto(ctx context.Context, req SetChatPhoto) (SetChatPhotoRes, error)
	// SetChatStickerSet implements setChatStickerSet operation.
	//
	// POST /setChatStickerSet
	SetChatStickerSet(ctx context.Context, req SetChatStickerSet) (SetChatStickerSetRes, error)
	// SetChatTitle implements setChatTitle operation.
	//
	// POST /setChatTitle
	SetChatTitle(ctx context.Context, req SetChatTitle) (SetChatTitleRes, error)
	// SetGameScore implements setGameScore operation.
	//
	// POST /setGameScore
	SetGameScore(ctx context.Context, req SetGameScore) (SetGameScoreRes, error)
	// SetMyCommands implements setMyCommands operation.
	//
	// POST /setMyCommands
	SetMyCommands(ctx context.Context, req SetMyCommands) (SetMyCommandsRes, error)
	// SetPassportDataErrors implements setPassportDataErrors operation.
	//
	// POST /setPassportDataErrors
	SetPassportDataErrors(ctx context.Context, req SetPassportDataErrors) (SetPassportDataErrorsRes, error)
	// SetStickerPositionInSet implements setStickerPositionInSet operation.
	//
	// POST /setStickerPositionInSet
	SetStickerPositionInSet(ctx context.Context, req SetStickerPositionInSet) (SetStickerPositionInSetRes, error)
	// SetStickerSetThumb implements setStickerSetThumb operation.
	//
	// POST /setStickerSetThumb
	SetStickerSetThumb(ctx context.Context, req SetStickerSetThumb) (SetStickerSetThumbRes, error)
	// SetWebhook implements setWebhook operation.
	//
	// POST /setWebhook
	SetWebhook(ctx context.Context, req SetWebhook) (SetWebhookRes, error)
	// StopMessageLiveLocation implements stopMessageLiveLocation operation.
	//
	// POST /stopMessageLiveLocation
	StopMessageLiveLocation(ctx context.Context, req StopMessageLiveLocation) (StopMessageLiveLocationRes, error)
	// StopPoll implements stopPoll operation.
	//
	// POST /stopPoll
	StopPoll(ctx context.Context, req StopPoll) (StopPollRes, error)
	// UnbanChatMember implements unbanChatMember operation.
	//
	// POST /unbanChatMember
	UnbanChatMember(ctx context.Context, req UnbanChatMember) (UnbanChatMemberRes, error)
	// UnpinAllChatMessages implements unpinAllChatMessages operation.
	//
	// POST /unpinAllChatMessages
	UnpinAllChatMessages(ctx context.Context, req UnpinAllChatMessages) (UnpinAllChatMessagesRes, error)
	// UnpinChatMessage implements unpinChatMessage operation.
	//
	// POST /unpinChatMessage
	UnpinChatMessage(ctx context.Context, req UnpinChatMessage) (UnpinChatMessageRes, error)
	// UploadStickerFile implements uploadStickerFile operation.
	//
	// POST /uploadStickerFile
	UploadStickerFile(ctx context.Context, req UploadStickerFile) (UploadStickerFileRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	cfg config
}

func NewServer(h Handler, opts ...Option) *Server {
	srv := &Server{
		h:   h,
		cfg: newConfig(opts...),
	}
	return srv
}

func (s *Server) Register(mux chi.Mux) {
	mux.MethodFunc("POST", "/addStickerToSet", s.HandleAddStickerToSetRequest)
	mux.MethodFunc("POST", "/answerCallbackQuery", s.HandleAnswerCallbackQueryRequest)
	mux.MethodFunc("POST", "/answerInlineQuery", s.HandleAnswerInlineQueryRequest)
	mux.MethodFunc("POST", "/answerPreCheckoutQuery", s.HandleAnswerPreCheckoutQueryRequest)
	mux.MethodFunc("POST", "/answerShippingQuery", s.HandleAnswerShippingQueryRequest)
	mux.MethodFunc("POST", "/banChatMember", s.HandleBanChatMemberRequest)
	mux.MethodFunc("POST", "/copyMessage", s.HandleCopyMessageRequest)
	mux.MethodFunc("POST", "/createChatInviteLink", s.HandleCreateChatInviteLinkRequest)
	mux.MethodFunc("POST", "/createNewStickerSet", s.HandleCreateNewStickerSetRequest)
	mux.MethodFunc("POST", "/deleteChatPhoto", s.HandleDeleteChatPhotoRequest)
	mux.MethodFunc("POST", "/deleteChatStickerSet", s.HandleDeleteChatStickerSetRequest)
	mux.MethodFunc("POST", "/deleteMessage", s.HandleDeleteMessageRequest)
	mux.MethodFunc("POST", "/deleteMyCommands", s.HandleDeleteMyCommandsRequest)
	mux.MethodFunc("POST", "/deleteStickerFromSet", s.HandleDeleteStickerFromSetRequest)
	mux.MethodFunc("POST", "/deleteWebhook", s.HandleDeleteWebhookRequest)
	mux.MethodFunc("POST", "/editChatInviteLink", s.HandleEditChatInviteLinkRequest)
	mux.MethodFunc("POST", "/editMessageCaption", s.HandleEditMessageCaptionRequest)
	mux.MethodFunc("POST", "/editMessageLiveLocation", s.HandleEditMessageLiveLocationRequest)
	mux.MethodFunc("POST", "/editMessageMedia", s.HandleEditMessageMediaRequest)
	mux.MethodFunc("POST", "/editMessageReplyMarkup", s.HandleEditMessageReplyMarkupRequest)
	mux.MethodFunc("POST", "/editMessageText", s.HandleEditMessageTextRequest)
	mux.MethodFunc("POST", "/exportChatInviteLink", s.HandleExportChatInviteLinkRequest)
	mux.MethodFunc("POST", "/forwardMessage", s.HandleForwardMessageRequest)
	mux.MethodFunc("POST", "/getChat", s.HandleGetChatRequest)
	mux.MethodFunc("POST", "/getChatAdministrators", s.HandleGetChatAdministratorsRequest)
	mux.MethodFunc("POST", "/getChatMember", s.HandleGetChatMemberRequest)
	mux.MethodFunc("POST", "/getChatMemberCount", s.HandleGetChatMemberCountRequest)
	mux.MethodFunc("POST", "/getFile", s.HandleGetFileRequest)
	mux.MethodFunc("POST", "/getGameHighScores", s.HandleGetGameHighScoresRequest)
	mux.MethodFunc("POST", "/getMe", s.HandleGetMeRequest)
	mux.MethodFunc("POST", "/getMyCommands", s.HandleGetMyCommandsRequest)
	mux.MethodFunc("POST", "/getStickerSet", s.HandleGetStickerSetRequest)
	mux.MethodFunc("POST", "/getUpdates", s.HandleGetUpdatesRequest)
	mux.MethodFunc("POST", "/getUserProfilePhotos", s.HandleGetUserProfilePhotosRequest)
	mux.MethodFunc("POST", "/leaveChat", s.HandleLeaveChatRequest)
	mux.MethodFunc("POST", "/pinChatMessage", s.HandlePinChatMessageRequest)
	mux.MethodFunc("POST", "/promoteChatMember", s.HandlePromoteChatMemberRequest)
	mux.MethodFunc("POST", "/restrictChatMember", s.HandleRestrictChatMemberRequest)
	mux.MethodFunc("POST", "/revokeChatInviteLink", s.HandleRevokeChatInviteLinkRequest)
	mux.MethodFunc("POST", "/sendAnimation", s.HandleSendAnimationRequest)
	mux.MethodFunc("POST", "/sendAudio", s.HandleSendAudioRequest)
	mux.MethodFunc("POST", "/sendChatAction", s.HandleSendChatActionRequest)
	mux.MethodFunc("POST", "/sendContact", s.HandleSendContactRequest)
	mux.MethodFunc("POST", "/sendDice", s.HandleSendDiceRequest)
	mux.MethodFunc("POST", "/sendDocument", s.HandleSendDocumentRequest)
	mux.MethodFunc("POST", "/sendGame", s.HandleSendGameRequest)
	mux.MethodFunc("POST", "/sendInvoice", s.HandleSendInvoiceRequest)
	mux.MethodFunc("POST", "/sendLocation", s.HandleSendLocationRequest)
	mux.MethodFunc("POST", "/sendMediaGroup", s.HandleSendMediaGroupRequest)
	mux.MethodFunc("POST", "/sendMessage", s.HandleSendMessageRequest)
	mux.MethodFunc("POST", "/sendPhoto", s.HandleSendPhotoRequest)
	mux.MethodFunc("POST", "/sendPoll", s.HandleSendPollRequest)
	mux.MethodFunc("POST", "/sendSticker", s.HandleSendStickerRequest)
	mux.MethodFunc("POST", "/sendVenue", s.HandleSendVenueRequest)
	mux.MethodFunc("POST", "/sendVideo", s.HandleSendVideoRequest)
	mux.MethodFunc("POST", "/sendVideoNote", s.HandleSendVideoNoteRequest)
	mux.MethodFunc("POST", "/sendVoice", s.HandleSendVoiceRequest)
	mux.MethodFunc("POST", "/setChatAdministratorCustomTitle", s.HandleSetChatAdministratorCustomTitleRequest)
	mux.MethodFunc("POST", "/setChatDescription", s.HandleSetChatDescriptionRequest)
	mux.MethodFunc("POST", "/setChatPermissions", s.HandleSetChatPermissionsRequest)
	mux.MethodFunc("POST", "/setChatPhoto", s.HandleSetChatPhotoRequest)
	mux.MethodFunc("POST", "/setChatStickerSet", s.HandleSetChatStickerSetRequest)
	mux.MethodFunc("POST", "/setChatTitle", s.HandleSetChatTitleRequest)
	mux.MethodFunc("POST", "/setGameScore", s.HandleSetGameScoreRequest)
	mux.MethodFunc("POST", "/setMyCommands", s.HandleSetMyCommandsRequest)
	mux.MethodFunc("POST", "/setPassportDataErrors", s.HandleSetPassportDataErrorsRequest)
	mux.MethodFunc("POST", "/setStickerPositionInSet", s.HandleSetStickerPositionInSetRequest)
	mux.MethodFunc("POST", "/setStickerSetThumb", s.HandleSetStickerSetThumbRequest)
	mux.MethodFunc("POST", "/setWebhook", s.HandleSetWebhookRequest)
	mux.MethodFunc("POST", "/stopMessageLiveLocation", s.HandleStopMessageLiveLocationRequest)
	mux.MethodFunc("POST", "/stopPoll", s.HandleStopPollRequest)
	mux.MethodFunc("POST", "/unbanChatMember", s.HandleUnbanChatMemberRequest)
	mux.MethodFunc("POST", "/unpinAllChatMessages", s.HandleUnpinAllChatMessagesRequest)
	mux.MethodFunc("POST", "/unpinChatMessage", s.HandleUnpinChatMessageRequest)
	mux.MethodFunc("POST", "/uploadStickerFile", s.HandleUploadStickerFileRequest)
}
