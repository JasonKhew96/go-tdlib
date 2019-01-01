package tdlib

import (
	"encoding/json"
	"fmt"
)

// GetAuthorizationState Returns the current authorization state; this is an offline request. For informational purposes only. Use updateAuthorizationState instead to maintain the current authorization state
func (client *Client) GetAuthorizationState() (AuthorizationState, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getAuthorizationState",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	switch AuthorizationStateEnum(result.Data["@type"].(string)) {

	case AuthorizationStateWaitTdlibParametersType:
		var authorizationState AuthorizationStateWaitTdlibParameters
		err = json.Unmarshal(result.Raw, &authorizationState)
		return &authorizationState, err

	case AuthorizationStateWaitEncryptionKeyType:
		var authorizationState AuthorizationStateWaitEncryptionKey
		err = json.Unmarshal(result.Raw, &authorizationState)
		return &authorizationState, err

	case AuthorizationStateWaitPhoneNumberType:
		var authorizationState AuthorizationStateWaitPhoneNumber
		err = json.Unmarshal(result.Raw, &authorizationState)
		return &authorizationState, err

	case AuthorizationStateWaitCodeType:
		var authorizationState AuthorizationStateWaitCode
		err = json.Unmarshal(result.Raw, &authorizationState)
		return &authorizationState, err

	case AuthorizationStateWaitPasswordType:
		var authorizationState AuthorizationStateWaitPassword
		err = json.Unmarshal(result.Raw, &authorizationState)
		return &authorizationState, err

	case AuthorizationStateReadyType:
		var authorizationState AuthorizationStateReady
		err = json.Unmarshal(result.Raw, &authorizationState)
		return &authorizationState, err

	case AuthorizationStateLoggingOutType:
		var authorizationState AuthorizationStateLoggingOut
		err = json.Unmarshal(result.Raw, &authorizationState)
		return &authorizationState, err

	case AuthorizationStateClosingType:
		var authorizationState AuthorizationStateClosing
		err = json.Unmarshal(result.Raw, &authorizationState)
		return &authorizationState, err

	case AuthorizationStateClosedType:
		var authorizationState AuthorizationStateClosed
		err = json.Unmarshal(result.Raw, &authorizationState)
		return &authorizationState, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// SetTdlibParameters Sets the parameters for TDLib initialization. Works only when the current authorization state is authorizationStateWaitTdlibParameters
// @param parameters Parameters
func (client *Client) SetTdlibParameters(parameters *TdlibParameters) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "setTdlibParameters",
		"parameters": parameters,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// CheckDatabaseEncryptionKey Checks the database encryption key for correctness. Works only when the current authorization state is authorizationStateWaitEncryptionKey
// @param encryptionKey Encryption key to check or set up
func (client *Client) CheckDatabaseEncryptionKey(encryptionKey []byte) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "checkDatabaseEncryptionKey",
		"encryption_key": encryptionKey,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetAuthenticationPhoneNumber Sets the phone number of the user and sends an authentication code to the user. Works only when the current authorization state is authorizationStateWaitPhoneNumber
// @param phoneNumber The phone number of the user, in international format
// @param allowFlashCall Pass true if the authentication code may be sent via flash call to the specified phone number
// @param isCurrentPhoneNumber Pass true if the phone number is used on the current device. Ignored if allow_flash_call is false
func (client *Client) SetAuthenticationPhoneNumber(phoneNumber string, allowFlashCall bool, isCurrentPhoneNumber bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                   "setAuthenticationPhoneNumber",
		"phone_number":            phoneNumber,
		"allow_flash_call":        allowFlashCall,
		"is_current_phone_number": isCurrentPhoneNumber,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ResendAuthenticationCode Re-sends an authentication code to the user. Works only when the current authorization state is authorizationStateWaitCode and the next_code_type of the result is not null
func (client *Client) ResendAuthenticationCode() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "resendAuthenticationCode",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// CheckAuthenticationCode Checks the authentication code. Works only when the current authorization state is authorizationStateWaitCode
// @param code The verification code received via SMS, Telegram message, phone call, or flash call
// @param firstName If the user is not yet registered, the first name of the user; 1-255 characters
// @param lastName If the user is not yet registered; the last name of the user; optional; 0-255 characters
func (client *Client) CheckAuthenticationCode(code string, firstName string, lastName string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "checkAuthenticationCode",
		"code":       code,
		"first_name": firstName,
		"last_name":  lastName,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// CheckAuthenticationPassword Checks the authentication password for correctness. Works only when the current authorization state is authorizationStateWaitPassword
// @param password The password to check
func (client *Client) CheckAuthenticationPassword(password string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "checkAuthenticationPassword",
		"password": password,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// RequestAuthenticationPasswordRecovery Requests to send a password recovery code to an email address that was previously set up. Works only when the current authorization state is authorizationStateWaitPassword
func (client *Client) RequestAuthenticationPasswordRecovery() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "requestAuthenticationPasswordRecovery",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// RecoverAuthenticationPassword Recovers the password with a password recovery code sent to an email address that was previously set up. Works only when the current authorization state is authorizationStateWaitPassword
// @param recoveryCode Recovery code to check
func (client *Client) RecoverAuthenticationPassword(recoveryCode string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "recoverAuthenticationPassword",
		"recovery_code": recoveryCode,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// CheckAuthenticationBotToken Checks the authentication token of a bot; to log in as a bot. Works only when the current authorization state is authorizationStateWaitPhoneNumber. Can be used instead of setAuthenticationPhoneNumber and checkAuthenticationCode to log in
// @param token The bot token
func (client *Client) CheckAuthenticationBotToken(token string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "checkAuthenticationBotToken",
		"token": token,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var okDummy Ok
	err = json.Unmarshal(result.Raw, &okDummy)
	return &okDummy, err

}

// LogOut Closes the TDLib instance after a proper logout. Requires an available network connection. All local data will be destroyed. After the logout completes, updateAuthorizationState with authorizationStateClosed will be sent
func (client *Client) LogOut() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "logOut",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// Close Closes the TDLib instance. All databases will be flushed to disk and properly closed. After the close completes, updateAuthorizationState with authorizationStateClosed will be sent
func (client *Client) Close() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "close",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// Destroy Closes the TDLib instance, destroying all local data without a proper logout. The current user session will remain in the list of all active sessions. All local data will be destroyed. After the destruction completes updateAuthorizationState with authorizationStateClosed will be sent
func (client *Client) Destroy() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "destroy",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetDatabaseEncryptionKey Changes the database encryption key. Usually the encryption key is never changed and is stored in some OS keychain
// @param newEncryptionKey New encryption key
func (client *Client) SetDatabaseEncryptionKey(newEncryptionKey []byte) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":              "setDatabaseEncryptionKey",
		"new_encryption_key": newEncryptionKey,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetPasswordState Returns the current state of 2-step verification
func (client *Client) GetPasswordState() (*PasswordState, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getPasswordState",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var passwordState PasswordState
	err = json.Unmarshal(result.Raw, &passwordState)
	return &passwordState, err

}

// SetPassword Changes the password for the user. If a new recovery email address is specified, then the error EMAIL_UNCONFIRMED is returned and the password change will not be applied until the new recovery email address has been confirmed. The application should periodically call getPasswordState to check whether the new email address has been confirmed
// @param oldPassword Previous password of the user
// @param newPassword New password of the user; may be empty to remove the password
// @param newHint New password hint; may be empty
// @param setRecoveryEmailAddress Pass true if the recovery email address should be changed
// @param newRecoveryEmailAddress New recovery email address; may be empty
func (client *Client) SetPassword(oldPassword string, newPassword string, newHint string, setRecoveryEmailAddress bool, newRecoveryEmailAddress string) (*PasswordState, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                      "setPassword",
		"old_password":               oldPassword,
		"new_password":               newPassword,
		"new_hint":                   newHint,
		"set_recovery_email_address": setRecoveryEmailAddress,
		"new_recovery_email_address": newRecoveryEmailAddress,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var passwordState PasswordState
	err = json.Unmarshal(result.Raw, &passwordState)
	return &passwordState, err

}

// GetRecoveryEmailAddress Returns a recovery email address that was previously set up. This method can be used to verify a password provided by the user
// @param password The password for the current user
func (client *Client) GetRecoveryEmailAddress(password string) (*RecoveryEmailAddress, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "getRecoveryEmailAddress",
		"password": password,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var recoveryEmailAddress RecoveryEmailAddress
	err = json.Unmarshal(result.Raw, &recoveryEmailAddress)
	return &recoveryEmailAddress, err

}

// SetRecoveryEmailAddress Changes the recovery email address of the user. If a new recovery email address is specified, then the error EMAIL_UNCONFIRMED is returned and the email address will not be changed until the new email has been confirmed. The application should periodically call getPasswordState to check whether the email address has been confirmed.
// @param password
// @param newRecoveryEmailAddress
func (client *Client) SetRecoveryEmailAddress(password string, newRecoveryEmailAddress string) (*PasswordState, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                      "setRecoveryEmailAddress",
		"password":                   password,
		"new_recovery_email_address": newRecoveryEmailAddress,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var passwordState PasswordState
	err = json.Unmarshal(result.Raw, &passwordState)
	return &passwordState, err

}

// RequestPasswordRecovery Requests to send a password recovery code to an email address that was previously set up
func (client *Client) RequestPasswordRecovery() (*EmailAddressAuthenticationCodeInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "requestPasswordRecovery",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var emailAddressAuthenticationCodeInfo EmailAddressAuthenticationCodeInfo
	err = json.Unmarshal(result.Raw, &emailAddressAuthenticationCodeInfo)
	return &emailAddressAuthenticationCodeInfo, err

}

// RecoverPassword Recovers the password using a recovery code sent to an email address that was previously set up
// @param recoveryCode Recovery code to check
func (client *Client) RecoverPassword(recoveryCode string) (*PasswordState, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "recoverPassword",
		"recovery_code": recoveryCode,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var passwordState PasswordState
	err = json.Unmarshal(result.Raw, &passwordState)
	return &passwordState, err

}

// CreateTemporaryPassword Creates a new temporary password for processing payments
// @param password Persistent user password
// @param validFor Time during which the temporary password will be valid, in seconds; should be between 60 and 86400
func (client *Client) CreateTemporaryPassword(password string, validFor int32) (*TemporaryPasswordState, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "createTemporaryPassword",
		"password":  password,
		"valid_for": validFor,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var temporaryPasswordState TemporaryPasswordState
	err = json.Unmarshal(result.Raw, &temporaryPasswordState)
	return &temporaryPasswordState, err

}

// GetTemporaryPasswordState Returns information about the current temporary password
func (client *Client) GetTemporaryPasswordState() (*TemporaryPasswordState, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getTemporaryPasswordState",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var temporaryPasswordState TemporaryPasswordState
	err = json.Unmarshal(result.Raw, &temporaryPasswordState)
	return &temporaryPasswordState, err

}

// ProcessDcUpdate Handles a DC_UPDATE push service notification. Can be called before authorization
// @param dc Value of the "dc" parameter of the notification
// @param addr Value of the "addr" parameter of the notification
func (client *Client) ProcessDcUpdate(dc string, addr string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "processDcUpdate",
		"dc":    dc,
		"addr":  addr,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetMe Returns the current user
func (client *Client) GetMe() (*User, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getMe",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var user User
	err = json.Unmarshal(result.Raw, &user)
	return &user, err

}

// GetUser Returns information about a user by their identifier. This is an offline request if the current user is not a bot
// @param userId User identifier
func (client *Client) GetUser(userId int32) (*User, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getUser",
		"user_id": userId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var userDummy User
	err = json.Unmarshal(result.Raw, &userDummy)
	return &userDummy, err

}

// GetUserFullInfo Returns full information about a user by their identifier
// @param userId User identifier
func (client *Client) GetUserFullInfo(userId int32) (*UserFullInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getUserFullInfo",
		"user_id": userId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var userFullInfo UserFullInfo
	err = json.Unmarshal(result.Raw, &userFullInfo)
	return &userFullInfo, err

}

// GetBasicGroup Returns information about a basic group by its identifier. This is an offline request if the current user is not a bot
// @param basicGroupId Basic group identifier
func (client *Client) GetBasicGroup(basicGroupId int32) (*BasicGroup, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "getBasicGroup",
		"basic_group_id": basicGroupId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var basicGroupDummy BasicGroup
	err = json.Unmarshal(result.Raw, &basicGroupDummy)
	return &basicGroupDummy, err

}

// GetBasicGroupFullInfo Returns full information about a basic group by its identifier
// @param basicGroupId Basic group identifier
func (client *Client) GetBasicGroupFullInfo(basicGroupId int32) (*BasicGroupFullInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "getBasicGroupFullInfo",
		"basic_group_id": basicGroupId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var basicGroupFullInfo BasicGroupFullInfo
	err = json.Unmarshal(result.Raw, &basicGroupFullInfo)
	return &basicGroupFullInfo, err

}

// GetSupergroup Returns information about a supergroup or channel by its identifier. This is an offline request if the current user is not a bot
// @param supergroupId Supergroup or channel identifier
func (client *Client) GetSupergroup(supergroupId int32) (*Supergroup, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "getSupergroup",
		"supergroup_id": supergroupId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var supergroupDummy Supergroup
	err = json.Unmarshal(result.Raw, &supergroupDummy)
	return &supergroupDummy, err

}

// GetSupergroupFullInfo Returns full information about a supergroup or channel by its identifier, cached for up to 1 minute
// @param supergroupId Supergroup or channel identifier
func (client *Client) GetSupergroupFullInfo(supergroupId int32) (*SupergroupFullInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "getSupergroupFullInfo",
		"supergroup_id": supergroupId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var supergroupFullInfo SupergroupFullInfo
	err = json.Unmarshal(result.Raw, &supergroupFullInfo)
	return &supergroupFullInfo, err

}

// GetSecretChat Returns information about a secret chat by its identifier. This is an offline request
// @param secretChatId Secret chat identifier
func (client *Client) GetSecretChat(secretChatId int32) (*SecretChat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "getSecretChat",
		"secret_chat_id": secretChatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var secretChatDummy SecretChat
	err = json.Unmarshal(result.Raw, &secretChatDummy)
	return &secretChatDummy, err

}

// GetChat Returns information about a chat by its identifier, this is an offline request if the current user is not a bot
// @param chatId Chat identifier
func (client *Client) GetChat(chatId int64) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getChat",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatDummy Chat
	err = json.Unmarshal(result.Raw, &chatDummy)
	return &chatDummy, err

}

// GetMessage Returns information about a message
// @param chatId Identifier of the chat the message belongs to
// @param messageId Identifier of the message to get
func (client *Client) GetMessage(chatId int64, messageId int64) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getMessage",
		"chat_id":    chatId,
		"message_id": messageId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// GetRepliedMessage Returns information about a message that is replied by given message
// @param chatId Identifier of the chat the message belongs to
// @param messageId Identifier of the message reply to which get
func (client *Client) GetRepliedMessage(chatId int64, messageId int64) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getRepliedMessage",
		"chat_id":    chatId,
		"message_id": messageId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// GetChatPinnedMessage Returns information about a pinned chat message
// @param chatId Identifier of the chat the message belongs to
func (client *Client) GetChatPinnedMessage(chatId int64) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getChatPinnedMessage",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var message Message
	err = json.Unmarshal(result.Raw, &message)
	return &message, err

}

// GetMessages Returns information about messages. If a message is not found, returns null on the corresponding position of the result
// @param chatId Identifier of the chat the messages belong to
// @param messageIds Identifiers of the messages to get
func (client *Client) GetMessages(chatId int64, messageIds []int64) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "getMessages",
		"chat_id":     chatId,
		"message_ids": messageIds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// GetFile Returns information about a file; this is an offline request
// @param fileId Identifier of the file to get
func (client *Client) GetFile(fileId int32) (*File, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getFile",
		"file_id": fileId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var fileDummy File
	err = json.Unmarshal(result.Raw, &fileDummy)
	return &fileDummy, err

}

// GetRemoteFile Returns information about a file by its remote ID; this is an offline request. Can be used to register a URL as a file for further uploading, or sending as a message
// @param remoteFileId Remote identifier of the file to get
// @param fileType File type, if known
func (client *Client) GetRemoteFile(remoteFileId string, fileType FileType) (*File, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "getRemoteFile",
		"remote_file_id": remoteFileId,
		"file_type":      fileType,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var fileDummy File
	err = json.Unmarshal(result.Raw, &fileDummy)
	return &fileDummy, err

}

// GetChats Returns an ordered list of chats. Chats are sorted by the pair (order, chat_id) in decreasing order. (For example, to get a list of chats from the beginning, the offset_order should be equal to 2^63 - 1).
// @param offsetOrder
// @param offsetChatId
// @param limit The maximum number of chats to be returned. It is possible that fewer chats than the limit are returned even if the end of the list is not reached
func (client *Client) GetChats(offsetOrder JSONInt64, offsetChatId int64, limit int32) (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "getChats",
		"offset_order":   offsetOrder,
		"offset_chat_id": offsetChatId,
		"limit":          limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}

// SearchPublicChat Searches a public chat by its username. Currently only private chats, supergroups and channels can be public. Returns the chat if found; otherwise an error is returned
// @param username Username to be resolved
func (client *Client) SearchPublicChat(username string) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "searchPublicChat",
		"username": username,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chat Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err

}

// SearchPublicChats Searches public chats by looking for specified query in their username and title. Currently only private chats, supergroups and channels can be public. Returns a meaningful number of results. Returns nothing if the length of the searched username prefix is less than 5. Excludes private chats with contacts and chats from the chat list from the results
// @param query Query to search for
func (client *Client) SearchPublicChats(query string) (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "searchPublicChats",
		"query": query,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}

// SearchChats Searches for the specified query in the title and username of already known chats, this is an offline request. Returns chats in the order seen in the chat list
// @param query Query to search for. If the query is empty, returns up to 20 recently found chats
// @param limit Maximum number of chats to be returned
func (client *Client) SearchChats(query string, limit int32) (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "searchChats",
		"query": query,
		"limit": limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}

// SearchChatsOnServer Searches for the specified query in the title and username of already known chats via request to the server. Returns chats in the order seen in the chat list
// @param query Query to search for
// @param limit Maximum number of chats to be returned
func (client *Client) SearchChatsOnServer(query string, limit int32) (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "searchChatsOnServer",
		"query": query,
		"limit": limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}

// GetTopChats Returns a list of frequently used chats. Supported only if the chat info database is enabled
// @param category Category of chats to be returned
// @param limit Maximum number of chats to be returned; up to 30
func (client *Client) GetTopChats(category TopChatCategory, limit int32) (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "getTopChats",
		"category": category,
		"limit":    limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}

// RemoveTopChat Removes a chat from the list of frequently used chats. Supported only if the chat info database is enabled
// @param category Category of frequently used chats
// @param chatId Chat identifier
func (client *Client) RemoveTopChat(category TopChatCategory, chatId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "removeTopChat",
		"category": category,
		"chat_id":  chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// AddRecentlyFoundChat Adds a chat to the list of recently found chats. The chat is added to the beginning of the list. If the chat is already in the list, it will be removed from the list first
// @param chatId Identifier of the chat to add
func (client *Client) AddRecentlyFoundChat(chatId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "addRecentlyFoundChat",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// RemoveRecentlyFoundChat Removes a chat from the list of recently found chats
// @param chatId Identifier of the chat to be removed
func (client *Client) RemoveRecentlyFoundChat(chatId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "removeRecentlyFoundChat",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ClearRecentlyFoundChats Clears the list of recently found chats
func (client *Client) ClearRecentlyFoundChats() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "clearRecentlyFoundChats",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// CheckChatUsername Checks whether a username can be set for a chat
// @param chatId Chat identifier; should be identifier of a supergroup chat, or a channel chat, or a private chat with self, or zero if chat is being created
// @param username Username to be checked
func (client *Client) CheckChatUsername(chatId JSONInt64, username string) (CheckChatUsernameResult, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "checkChatUsername",
		"chat_id":  chatId,
		"username": username,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	switch CheckChatUsernameResultEnum(result.Data["@type"].(string)) {

	case CheckChatUsernameResultOkType:
		var checkChatUsernameResult CheckChatUsernameResultOk
		err = json.Unmarshal(result.Raw, &checkChatUsernameResult)
		return &checkChatUsernameResult, err

	case CheckChatUsernameResultUsernameInvalidType:
		var checkChatUsernameResult CheckChatUsernameResultUsernameInvalid
		err = json.Unmarshal(result.Raw, &checkChatUsernameResult)
		return &checkChatUsernameResult, err

	case CheckChatUsernameResultUsernameOccupiedType:
		var checkChatUsernameResult CheckChatUsernameResultUsernameOccupied
		err = json.Unmarshal(result.Raw, &checkChatUsernameResult)
		return &checkChatUsernameResult, err

	case CheckChatUsernameResultPublicChatsTooMuchType:
		var checkChatUsernameResult CheckChatUsernameResultPublicChatsTooMuch
		err = json.Unmarshal(result.Raw, &checkChatUsernameResult)
		return &checkChatUsernameResult, err

	case CheckChatUsernameResultPublicGroupsUnavailableType:
		var checkChatUsernameResult CheckChatUsernameResultPublicGroupsUnavailable
		err = json.Unmarshal(result.Raw, &checkChatUsernameResult)
		return &checkChatUsernameResult, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// GetCreatedPublicChats Returns a list of public chats created by the user
func (client *Client) GetCreatedPublicChats() (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getCreatedPublicChats",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}

// GetGroupsInCommon Returns a list of common chats with a given user. Chats are sorted by their type and creation date
// @param userId User identifier
// @param offsetChatId Chat identifier starting from which to return chats; use 0 for the first request
// @param limit Maximum number of chats to be returned; up to 100
func (client *Client) GetGroupsInCommon(userId int32, offsetChatId int64, limit int32) (*Chats, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "getGroupsInCommon",
		"user_id":        userId,
		"offset_chat_id": offsetChatId,
		"limit":          limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chats Chats
	err = json.Unmarshal(result.Raw, &chats)
	return &chats, err

}

// GetChatHistory Returns messages in a chat. The messages are returned in a reverse chronological order (i.e., in order of decreasing message_id).
// @param chatId Chat identifier
// @param fromMessageId Identifier of the message starting from which history must be fetched; use 0 to get results from the last message
// @param offset Specify 0 to get results from exactly the from_message_id or a negative offset to get the specified message and some newer messages
// @param limit The maximum number of messages to be returned; must be positive and can't be greater than 100. If the offset is negative, the limit must be greater than -offset. Fewer messages may be returned than specified by the limit, even if the end of the message history has not been reached
// @param onlyLocal If true, returns only messages that are available locally without sending network requests
func (client *Client) GetChatHistory(chatId int64, fromMessageId int64, offset int32, limit int32, onlyLocal bool) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "getChatHistory",
		"chat_id":         chatId,
		"from_message_id": fromMessageId,
		"offset":          offset,
		"limit":           limit,
		"only_local":      onlyLocal,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// DeleteChatHistory Deletes all messages in the chat only for the user. Cannot be used in channels and public supergroups
// @param chatId Chat identifier
// @param removeFromChatList Pass true if the chat should be removed from the chats list
func (client *Client) DeleteChatHistory(chatId int64, removeFromChatList bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "deleteChatHistory",
		"chat_id":               chatId,
		"remove_from_chat_list": removeFromChatList,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SearchChatMessages Searches for messages with given words in the chat. Returns the results in reverse chronological order, i.e. in order of decreasing message_id. Cannot be used in secret chats with a non-empty query
// @param chatId Identifier of the chat in which to search messages
// @param query Query to search for
// @param senderUserId If not 0, only messages sent by the specified user will be returned. Not supported in secret chats
// @param fromMessageId Identifier of the message starting from which history must be fetched; use 0 to get results from the last message
// @param offset Specify 0 to get results from exactly the from_message_id or a negative offset to get the specified message and some newer messages
// @param limit The maximum number of messages to be returned; must be positive and can't be greater than 100. If the offset is negative, the limit must be greater than -offset. Fewer messages may be returned than specified by the limit, even if the end of the message history has not been reached
// @param filter Filter for message content in the search results
func (client *Client) SearchChatMessages(chatId int64, query string, senderUserId int32, fromMessageId int64, offset int32, limit int32, filter SearchMessagesFilter) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "searchChatMessages",
		"chat_id":         chatId,
		"query":           query,
		"sender_user_id":  senderUserId,
		"from_message_id": fromMessageId,
		"offset":          offset,
		"limit":           limit,
		"filter":          filter,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// SearchMessages Searches for messages in all chats except secret chats. Returns the results in reverse chronological order (i.e., in order of decreasing (date, chat_id, message_id)).
// @param query Query to search for
// @param offsetDate The date of the message starting from which the results should be fetched. Use 0 or any date in the future to get results from the last message
// @param offsetChatId The chat identifier of the last found message, or 0 for the first request
// @param offsetMessageId The message identifier of the last found message, or 0 for the first request
// @param limit The maximum number of messages to be returned, up to 100. Fewer messages may be returned than specified by the limit, even if the end of the message history has not been reached
func (client *Client) SearchMessages(query string, offsetDate int32, offsetChatId int64, offsetMessageId int64, limit int32) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "searchMessages",
		"query":             query,
		"offset_date":       offsetDate,
		"offset_chat_id":    offsetChatId,
		"offset_message_id": offsetMessageId,
		"limit":             limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// SearchSecretMessages Searches for messages in secret chats. Returns the results in reverse chronological order. For optimal performance the number of returned messages is chosen by the library
// @param chatId Identifier of the chat in which to search. Specify 0 to search in all secret chats
// @param query Query to search for. If empty, searchChatMessages should be used instead
// @param fromSearchId The identifier from the result of a previous request, use 0 to get results from the last message
// @param limit Maximum number of messages to be returned; up to 100. Fewer messages may be returned than specified by the limit, even if the end of the message history has not been reached
// @param filter A filter for the content of messages in the search results
func (client *Client) SearchSecretMessages(chatId int64, query string, fromSearchId JSONInt64, limit int32, filter SearchMessagesFilter) (*FoundMessages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "searchSecretMessages",
		"chat_id":        chatId,
		"query":          query,
		"from_search_id": fromSearchId,
		"limit":          limit,
		"filter":         filter,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var foundMessages FoundMessages
	err = json.Unmarshal(result.Raw, &foundMessages)
	return &foundMessages, err

}

// SearchCallMessages Searches for call messages. Returns the results in reverse chronological order (i. e., in order of decreasing message_id). For optimal performance the number of returned messages is chosen by the library
// @param fromMessageId Identifier of the message from which to search; use 0 to get results from the last message
// @param limit The maximum number of messages to be returned; up to 100. Fewer messages may be returned than specified by the limit, even if the end of the message history has not been reached
// @param onlyMissed If true, returns only messages with missed calls
func (client *Client) SearchCallMessages(fromMessageId int64, limit int32, onlyMissed bool) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "searchCallMessages",
		"from_message_id": fromMessageId,
		"limit":           limit,
		"only_missed":     onlyMissed,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// SearchChatRecentLocationMessages Returns information about the recent locations of chat members that were sent to the chat. Returns up to 1 location message per user
// @param chatId Chat identifier
// @param limit Maximum number of messages to be returned
func (client *Client) SearchChatRecentLocationMessages(chatId int64, limit int32) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "searchChatRecentLocationMessages",
		"chat_id": chatId,
		"limit":   limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// GetActiveLiveLocationMessages Returns all active live locations that should be updated by the client. The list is persistent across application restarts only if the message database is used
func (client *Client) GetActiveLiveLocationMessages() (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getActiveLiveLocationMessages",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// GetChatMessageByDate Returns the last message sent in a chat no later than the specified date
// @param chatId Chat identifier
// @param date Point in time (Unix timestamp) relative to which to search for messages
func (client *Client) GetChatMessageByDate(chatId int64, date int32) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getChatMessageByDate",
		"chat_id": chatId,
		"date":    date,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var message Message
	err = json.Unmarshal(result.Raw, &message)
	return &message, err

}

// GetChatMessageCount Returns approximate number of messages of the specified type in the chat
// @param chatId Identifier of the chat in which to count messages
// @param filter Filter for message content; searchMessagesFilterEmpty is unsupported in this function
// @param returnLocal If true, returns count that is available locally without sending network requests, returning -1 if the number of messages is unknown
func (client *Client) GetChatMessageCount(chatId int64, filter SearchMessagesFilter, returnLocal bool) (*Count, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "getChatMessageCount",
		"chat_id":      chatId,
		"filter":       filter,
		"return_local": returnLocal,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var count Count
	err = json.Unmarshal(result.Raw, &count)
	return &count, err

}

// GetPublicMessageLink Returns a public HTTPS link to a message. Available only for messages in public supergroups and channels
// @param chatId Identifier of the chat to which the message belongs
// @param messageId Identifier of the message
// @param forAlbum Pass true if a link for a whole media album should be returned
func (client *Client) GetPublicMessageLink(chatId int64, messageId int64, forAlbum bool) (*PublicMessageLink, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getPublicMessageLink",
		"chat_id":    chatId,
		"message_id": messageId,
		"for_album":  forAlbum,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var publicMessageLink PublicMessageLink
	err = json.Unmarshal(result.Raw, &publicMessageLink)
	return &publicMessageLink, err

}

// SendMessage Sends a message. Returns the sent message
// @param chatId Target chat
// @param replyToMessageId Identifier of the message to reply to or 0
// @param disableNotification Pass true to disable notification for the message. Not supported in secret chats
// @param fromBackground Pass true if the message is sent from the background
// @param replyMarkup Markup for replying to the message; for bots only
// @param inputMessageContent The content of the message to be sent
func (client *Client) SendMessage(chatId int64, replyToMessageId int64, disableNotification bool, fromBackground bool, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "sendMessage",
		"chat_id":               chatId,
		"reply_to_message_id":   replyToMessageId,
		"disable_notification":  disableNotification,
		"from_background":       fromBackground,
		"reply_markup":          replyMarkup,
		"input_message_content": inputMessageContent,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// SendMessageAlbum Sends messages grouped together into an album. Currently only photo and video messages can be grouped into an album. Returns sent messages
// @param chatId Target chat
// @param replyToMessageId Identifier of a message to reply to or 0
// @param disableNotification Pass true to disable notification for the messages. Not supported in secret chats
// @param fromBackground Pass true if the messages are sent from the background
// @param inputMessageContents Contents of messages to be sent
func (client *Client) SendMessageAlbum(chatId int64, replyToMessageId int64, disableNotification bool, fromBackground bool, inputMessageContents []InputMessageContent) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                  "sendMessageAlbum",
		"chat_id":                chatId,
		"reply_to_message_id":    replyToMessageId,
		"disable_notification":   disableNotification,
		"from_background":        fromBackground,
		"input_message_contents": inputMessageContents,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// SendBotStartMessage Invites a bot to a chat (if it is not yet a member) and sends it the /start command. Bots can't be invited to a private chat other than the chat with the bot. Bots can't be invited to channels (although they can be added as admins) and secret chats. Returns the sent message
// @param botUserId Identifier of the bot
// @param chatId Identifier of the target chat
// @param parameter A hidden parameter sent to the bot for deep linking purposes (https://api.telegram.org/bots#deep-linking)
func (client *Client) SendBotStartMessage(botUserId int32, chatId int64, parameter string) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "sendBotStartMessage",
		"bot_user_id": botUserId,
		"chat_id":     chatId,
		"parameter":   parameter,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var message Message
	err = json.Unmarshal(result.Raw, &message)
	return &message, err

}

// SendInlineQueryResultMessage Sends the result of an inline query as a message. Returns the sent message. Always clears a chat draft message
// @param chatId Target chat
// @param replyToMessageId Identifier of a message to reply to or 0
// @param disableNotification Pass true to disable notification for the message. Not supported in secret chats
// @param fromBackground Pass true if the message is sent from background
// @param queryId Identifier of the inline query
// @param resultId Identifier of the inline result
func (client *Client) SendInlineQueryResultMessage(chatId int64, replyToMessageId int64, disableNotification bool, fromBackground bool, queryId JSONInt64, resultId string) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                "sendInlineQueryResultMessage",
		"chat_id":              chatId,
		"reply_to_message_id":  replyToMessageId,
		"disable_notification": disableNotification,
		"from_background":      fromBackground,
		"query_id":             queryId,
		"result_id":            resultId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// ForwardMessages Forwards previously sent messages. Returns the forwarded messages in the same order as the message identifiers passed in message_ids. If a message can't be forwarded, null will be returned instead of the message
// @param chatId Identifier of the chat to which to forward messages
// @param fromChatId Identifier of the chat from which to forward messages
// @param messageIds Identifiers of the messages to forward
// @param disableNotification Pass true to disable notification for the message, doesn't work if messages are forwarded to a secret chat
// @param fromBackground Pass true if the message is sent from the background
// @param asAlbum True, if the messages should be grouped into an album after forwarding. For this to work, no more than 10 messages may be forwarded, and all of them must be photo or video messages
func (client *Client) ForwardMessages(chatId int64, fromChatId int64, messageIds []int64, disableNotification bool, fromBackground bool, asAlbum bool) (*Messages, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                "forwardMessages",
		"chat_id":              chatId,
		"from_chat_id":         fromChatId,
		"message_ids":          messageIds,
		"disable_notification": disableNotification,
		"from_background":      fromBackground,
		"as_album":             asAlbum,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messages Messages
	err = json.Unmarshal(result.Raw, &messages)
	return &messages, err

}

// SendChatSetTtlMessage Changes the current TTL setting (sets a new self-destruct timer) in a secret chat and sends the corresponding message
// @param chatId Chat identifier
// @param ttl New TTL value, in seconds
func (client *Client) SendChatSetTtlMessage(chatId int64, ttl int32) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "sendChatSetTtlMessage",
		"chat_id": chatId,
		"ttl":     ttl,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var message Message
	err = json.Unmarshal(result.Raw, &message)
	return &message, err

}

// SendChatScreenshotTakenNotification Sends a notification about a screenshot taken in a chat. Supported only in private and secret chats
// @param chatId Chat identifier
func (client *Client) SendChatScreenshotTakenNotification(chatId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "sendChatScreenshotTakenNotification",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// AddLocalMessage Adds a local message to a chat. The message is persistent across application restarts only if the message database is used. Returns the added message
// @param chatId Target chat
// @param senderUserId Identifier of the user who will be shown as the sender of the message; may be 0 for channel posts
// @param replyToMessageId Identifier of the message to reply to or 0
// @param disableNotification Pass true to disable notification for the message
// @param inputMessageContent The content of the message to be added
func (client *Client) AddLocalMessage(chatId int64, senderUserId int32, replyToMessageId int64, disableNotification bool, inputMessageContent InputMessageContent) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "addLocalMessage",
		"chat_id":               chatId,
		"sender_user_id":        senderUserId,
		"reply_to_message_id":   replyToMessageId,
		"disable_notification":  disableNotification,
		"input_message_content": inputMessageContent,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// DeleteMessages Deletes messages
// @param chatId Chat identifier
// @param messageIds Identifiers of the messages to be deleted
// @param revoke Pass true to try to delete outgoing messages for all chat members (may fail if messages are too old). Always true for supergroups, channels and secret chats
func (client *Client) DeleteMessages(chatId int64, messageIds []int64, revoke bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "deleteMessages",
		"chat_id":     chatId,
		"message_ids": messageIds,
		"revoke":      revoke,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var okDummy Ok
	err = json.Unmarshal(result.Raw, &okDummy)
	return &okDummy, err

}

// DeleteChatMessagesFromUser Deletes all messages sent by the specified user to a chat. Supported only in supergroups; requires can_delete_messages administrator privileges
// @param chatId Chat identifier
// @param userId User identifier
func (client *Client) DeleteChatMessagesFromUser(chatId int64, userId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "deleteChatMessagesFromUser",
		"chat_id": chatId,
		"user_id": userId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// EditMessageText Edits the text of a message (or a text of a game message). Returns the edited message after the edit is completed on the server side
// @param chatId The chat the message belongs to
// @param messageId Identifier of the message
// @param replyMarkup The new message reply markup; for bots only
// @param inputMessageContent New text content of the message. Should be of type InputMessageText
func (client *Client) EditMessageText(chatId int64, messageId int64, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "editMessageText",
		"chat_id":               chatId,
		"message_id":            messageId,
		"reply_markup":          replyMarkup,
		"input_message_content": inputMessageContent,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// EditMessageLiveLocation Edits the message content of a live location. Messages can be edited for a limited period of time specified in the live location. Returns the edited message after the edit is completed on the server side
// @param chatId The chat the message belongs to
// @param messageId Identifier of the message
// @param replyMarkup The new message reply markup; for bots only
// @param location New location content of the message; may be null. Pass null to stop sharing the live location
func (client *Client) EditMessageLiveLocation(chatId int64, messageId int64, replyMarkup ReplyMarkup, location *Location) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "editMessageLiveLocation",
		"chat_id":      chatId,
		"message_id":   messageId,
		"reply_markup": replyMarkup,
		"location":     location,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// EditMessageMedia Edits the content of a message with an animation, an audio, a document, a photo or a video. The media in the message can't be replaced if the message was set to self-destruct. Media can't be replaced by self-destructing media. Media in an album can be edited only to contain a photo or a video. Returns the edited message after the edit is completed on the server side
// @param chatId The chat the message belongs to
// @param messageId Identifier of the message
// @param replyMarkup The new message reply markup; for bots only
// @param inputMessageContent New content of the message. Must be one of the following types: InputMessageAnimation, InputMessageAudio, InputMessageDocument, InputMessagePhoto or InputMessageVideo
func (client *Client) EditMessageMedia(chatId int64, messageId int64, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "editMessageMedia",
		"chat_id":               chatId,
		"message_id":            messageId,
		"reply_markup":          replyMarkup,
		"input_message_content": inputMessageContent,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// EditMessageCaption Edits the message content caption. Returns the edited message after the edit is completed on the server side
// @param chatId The chat the message belongs to
// @param messageId Identifier of the message
// @param replyMarkup The new message reply markup; for bots only
// @param caption New message content caption; 0-GetOption("message_caption_length_max") characters
func (client *Client) EditMessageCaption(chatId int64, messageId int64, replyMarkup ReplyMarkup, caption *FormattedText) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "editMessageCaption",
		"chat_id":      chatId,
		"message_id":   messageId,
		"reply_markup": replyMarkup,
		"caption":      caption,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// EditMessageReplyMarkup Edits the message reply markup; for bots only. Returns the edited message after the edit is completed on the server side
// @param chatId The chat the message belongs to
// @param messageId Identifier of the message
// @param replyMarkup The new message reply markup
func (client *Client) EditMessageReplyMarkup(chatId int64, messageId int64, replyMarkup ReplyMarkup) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "editMessageReplyMarkup",
		"chat_id":      chatId,
		"message_id":   messageId,
		"reply_markup": replyMarkup,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// EditInlineMessageText Edits the text of an inline text or game message sent via a bot; for bots only
// @param inlineMessageId Inline message identifier
// @param replyMarkup The new message reply markup
// @param inputMessageContent New text content of the message. Should be of type InputMessageText
func (client *Client) EditInlineMessageText(inlineMessageId string, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "editInlineMessageText",
		"inline_message_id":     inlineMessageId,
		"reply_markup":          replyMarkup,
		"input_message_content": inputMessageContent,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// EditInlineMessageLiveLocation Edits the content of a live location in an inline message sent via a bot; for bots only
// @param inlineMessageId Inline message identifier
// @param replyMarkup The new message reply markup
// @param location New location content of the message; may be null. Pass null to stop sharing the live location
func (client *Client) EditInlineMessageLiveLocation(inlineMessageId string, replyMarkup ReplyMarkup, location *Location) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "editInlineMessageLiveLocation",
		"inline_message_id": inlineMessageId,
		"reply_markup":      replyMarkup,
		"location":          location,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// EditInlineMessageMedia Edits the content of a message with an animation, an audio, a document, a photo or a video in an inline message sent via a bot; for bots only
// @param inlineMessageId Inline message identifier
// @param replyMarkup The new message reply markup; for bots only
// @param inputMessageContent New content of the message. Must be one of the following types: InputMessageAnimation, InputMessageAudio, InputMessageDocument, InputMessagePhoto or InputMessageVideo
func (client *Client) EditInlineMessageMedia(inlineMessageId string, replyMarkup ReplyMarkup, inputMessageContent InputMessageContent) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "editInlineMessageMedia",
		"inline_message_id":     inlineMessageId,
		"reply_markup":          replyMarkup,
		"input_message_content": inputMessageContent,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// EditInlineMessageCaption Edits the caption of an inline message sent via a bot; for bots only
// @param inlineMessageId Inline message identifier
// @param replyMarkup The new message reply markup
// @param caption New message content caption; 0-GetOption("message_caption_length_max") characters
func (client *Client) EditInlineMessageCaption(inlineMessageId string, replyMarkup ReplyMarkup, caption *FormattedText) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "editInlineMessageCaption",
		"inline_message_id": inlineMessageId,
		"reply_markup":      replyMarkup,
		"caption":           caption,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// EditInlineMessageReplyMarkup Edits the reply markup of an inline message sent via a bot; for bots only
// @param inlineMessageId Inline message identifier
// @param replyMarkup The new message reply markup
func (client *Client) EditInlineMessageReplyMarkup(inlineMessageId string, replyMarkup ReplyMarkup) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "editInlineMessageReplyMarkup",
		"inline_message_id": inlineMessageId,
		"reply_markup":      replyMarkup,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetTextEntities Returns all entities (mentions, hashtags, cashtags, bot commands, URLs, and email addresses) contained in the text. This is an offline method. Can be called before authorization. Can be called synchronously
// @param text The text in which to look for entites
func (client *Client) GetTextEntities(text string) (*TextEntities, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getTextEntities",
		"text":  text,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var textEntities TextEntities
	err = json.Unmarshal(result.Raw, &textEntities)
	return &textEntities, err

}

// ParseTextEntities Parses Bold, Italic, Code, Pre, PreCode and TextUrl entities contained in the text. This is an offline method. Can be called before authorization. Can be called synchronously
// @param text The text which should be parsed
// @param parseMode Text parse mode
func (client *Client) ParseTextEntities(text string, parseMode TextParseMode) (*FormattedText, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "parseTextEntities",
		"text":       text,
		"parse_mode": parseMode,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var formattedText FormattedText
	err = json.Unmarshal(result.Raw, &formattedText)
	return &formattedText, err

}

// GetFileMimeType Returns the MIME type of a file, guessed by its extension. Returns an empty string on failure. This is an offline method. Can be called before authorization. Can be called synchronously
// @param fileName The name of the file or path to the file
func (client *Client) GetFileMimeType(fileName string) (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "getFileMimeType",
		"file_name": fileName,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// GetFileExtension Returns the extension of a file, guessed by its MIME type. Returns an empty string on failure. This is an offline method. Can be called before authorization. Can be called synchronously
// @param mimeType The MIME type of the file
func (client *Client) GetFileExtension(mimeType string) (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "getFileExtension",
		"mime_type": mimeType,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// CleanFileName Removes potentially dangerous characters from the name of a file. The encoding of the file name is supposed to be UTF-8. Returns an empty string on failure. This is an offline method. Can be called before authorization. Can be called synchronously
// @param fileName File name or path to the file
func (client *Client) CleanFileName(fileName string) (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "cleanFileName",
		"file_name": fileName,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// GetLanguagePackString Returns a string stored in the local database from the specified localization target and language pack by its key. Returns a 404 error if the string is not found. This is an offline method. Can be called before authorization. Can be called synchronously
// @param languagePackDatabasePath Path to the language pack database in which strings are stored
// @param localizationTarget Localization target to which the language pack belongs
// @param languagePackId Language pack identifier
// @param key Language pack key of the string to be returned
func (client *Client) GetLanguagePackString(languagePackDatabasePath string, localizationTarget string, languagePackId string, key string) (LanguagePackStringValue, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                       "getLanguagePackString",
		"language_pack_database_path": languagePackDatabasePath,
		"localization_target":         localizationTarget,
		"language_pack_id":            languagePackId,
		"key":                         key,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	switch LanguagePackStringValueEnum(result.Data["@type"].(string)) {

	case LanguagePackStringValueOrdinaryType:
		var languagePackStringValue LanguagePackStringValueOrdinary
		err = json.Unmarshal(result.Raw, &languagePackStringValue)
		return &languagePackStringValue, err

	case LanguagePackStringValuePluralizedType:
		var languagePackStringValue LanguagePackStringValuePluralized
		err = json.Unmarshal(result.Raw, &languagePackStringValue)
		return &languagePackStringValue, err

	case LanguagePackStringValueDeletedType:
		var languagePackStringValue LanguagePackStringValueDeleted
		err = json.Unmarshal(result.Raw, &languagePackStringValue)
		return &languagePackStringValue, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// GetInlineQueryResults Sends an inline query to a bot and returns its results. Returns an error with code 502 if the bot fails to answer the query before the query timeout expires
// @param botUserId The identifier of the target bot
// @param chatId Identifier of the chat, where the query was sent
// @param userLocation Location of the user, only if needed
// @param query Text of the query
// @param offset Offset of the first entry to return
func (client *Client) GetInlineQueryResults(botUserId int32, chatId int64, userLocation *Location, query string, offset string) (*InlineQueryResults, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "getInlineQueryResults",
		"bot_user_id":   botUserId,
		"chat_id":       chatId,
		"user_location": userLocation,
		"query":         query,
		"offset":        offset,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var inlineQueryResults InlineQueryResults
	err = json.Unmarshal(result.Raw, &inlineQueryResults)
	return &inlineQueryResults, err

}

// AnswerInlineQuery Sets the result of an inline query; for bots only
// @param inlineQueryId Identifier of the inline query
// @param isPersonal True, if the result of the query can be cached for the specified user
// @param results The results of the query
// @param cacheTime Allowed time to cache the results of the query, in seconds
// @param nextOffset Offset for the next inline query; pass an empty string if there are no more results
// @param switchPmText If non-empty, this text should be shown on the button that opens a private chat with the bot and sends a start message to the bot with the parameter switch_pm_parameter
// @param switchPmParameter The parameter for the bot start message
func (client *Client) AnswerInlineQuery(inlineQueryId JSONInt64, isPersonal bool, results []InputInlineQueryResult, cacheTime int32, nextOffset string, switchPmText string, switchPmParameter string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":               "answerInlineQuery",
		"inline_query_id":     inlineQueryId,
		"is_personal":         isPersonal,
		"results":             results,
		"cache_time":          cacheTime,
		"next_offset":         nextOffset,
		"switch_pm_text":      switchPmText,
		"switch_pm_parameter": switchPmParameter,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetCallbackQueryAnswer Sends a callback query to a bot and returns an answer. Returns an error with code 502 if the bot fails to answer the query before the query timeout expires
// @param chatId Identifier of the chat with the message
// @param messageId Identifier of the message from which the query originated
// @param payload Query payload
func (client *Client) GetCallbackQueryAnswer(chatId int64, messageId int64, payload CallbackQueryPayload) (*CallbackQueryAnswer, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getCallbackQueryAnswer",
		"chat_id":    chatId,
		"message_id": messageId,
		"payload":    payload,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var callbackQueryAnswer CallbackQueryAnswer
	err = json.Unmarshal(result.Raw, &callbackQueryAnswer)
	return &callbackQueryAnswer, err

}

// AnswerCallbackQuery Sets the result of a callback query; for bots only
// @param callbackQueryId Identifier of the callback query
// @param text Text of the answer
// @param showAlert If true, an alert should be shown to the user instead of a toast notification
// @param url URL to be opened
// @param cacheTime Time during which the result of the query can be cached, in seconds
func (client *Client) AnswerCallbackQuery(callbackQueryId JSONInt64, text string, showAlert bool, url string, cacheTime int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "answerCallbackQuery",
		"callback_query_id": callbackQueryId,
		"text":              text,
		"show_alert":        showAlert,
		"url":               url,
		"cache_time":        cacheTime,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// AnswerShippingQuery Sets the result of a shipping query; for bots only
// @param shippingQueryId Identifier of the shipping query
// @param shippingOptions Available shipping options
// @param errorMessage An error message, empty on success
func (client *Client) AnswerShippingQuery(shippingQueryId JSONInt64, shippingOptions []ShippingOption, errorMessage string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "answerShippingQuery",
		"shipping_query_id": shippingQueryId,
		"shipping_options":  shippingOptions,
		"error_message":     errorMessage,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// AnswerPreCheckoutQuery Sets the result of a pre-checkout query; for bots only
// @param preCheckoutQueryId Identifier of the pre-checkout query
// @param errorMessage An error message, empty on success
func (client *Client) AnswerPreCheckoutQuery(preCheckoutQueryId JSONInt64, errorMessage string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "answerPreCheckoutQuery",
		"pre_checkout_query_id": preCheckoutQueryId,
		"error_message":         errorMessage,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetGameScore Updates the game score of the specified user in the game; for bots only
// @param chatId The chat to which the message with the game
// @param messageId Identifier of the message
// @param editMessage True, if the message should be edited
// @param userId User identifier
// @param score The new score
// @param force Pass true to update the score even if it decreases. If the score is 0, the user will be deleted from the high score table
func (client *Client) SetGameScore(chatId int64, messageId int64, editMessage bool, userId int32, score int32, force bool) (*Message, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "setGameScore",
		"chat_id":      chatId,
		"message_id":   messageId,
		"edit_message": editMessage,
		"user_id":      userId,
		"score":        score,
		"force":        force,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var messageDummy Message
	err = json.Unmarshal(result.Raw, &messageDummy)
	return &messageDummy, err

}

// SetInlineGameScore Updates the game score of the specified user in a game; for bots only
// @param inlineMessageId Inline message identifier
// @param editMessage True, if the message should be edited
// @param userId User identifier
// @param score The new score
// @param force Pass true to update the score even if it decreases. If the score is 0, the user will be deleted from the high score table
func (client *Client) SetInlineGameScore(inlineMessageId string, editMessage bool, userId int32, score int32, force bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "setInlineGameScore",
		"inline_message_id": inlineMessageId,
		"edit_message":      editMessage,
		"user_id":           userId,
		"score":             score,
		"force":             force,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetGameHighScores Returns the high scores for a game and some part of the high score table in the range of the specified user; for bots only
// @param chatId The chat that contains the message with the game
// @param messageId Identifier of the message
// @param userId User identifier
func (client *Client) GetGameHighScores(chatId int64, messageId int64, userId int32) (*GameHighScores, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getGameHighScores",
		"chat_id":    chatId,
		"message_id": messageId,
		"user_id":    userId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var gameHighScores GameHighScores
	err = json.Unmarshal(result.Raw, &gameHighScores)
	return &gameHighScores, err

}

// GetInlineGameHighScores Returns game high scores and some part of the high score table in the range of the specified user; for bots only
// @param inlineMessageId Inline message identifier
// @param userId User identifier
func (client *Client) GetInlineGameHighScores(inlineMessageId string, userId int32) (*GameHighScores, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "getInlineGameHighScores",
		"inline_message_id": inlineMessageId,
		"user_id":           userId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var gameHighScores GameHighScores
	err = json.Unmarshal(result.Raw, &gameHighScores)
	return &gameHighScores, err

}

// DeleteChatReplyMarkup Deletes the default reply markup from a chat. Must be called after a one-time keyboard or a ForceReply reply markup has been used. UpdateChatReplyMarkup will be sent if the reply markup will be changed
// @param chatId Chat identifier
// @param messageId The message identifier of the used keyboard
func (client *Client) DeleteChatReplyMarkup(chatId int64, messageId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "deleteChatReplyMarkup",
		"chat_id":    chatId,
		"message_id": messageId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SendChatAction Sends a notification about user activity in a chat
// @param chatId Chat identifier
// @param action The action description
func (client *Client) SendChatAction(chatId int64, action ChatAction) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "sendChatAction",
		"chat_id": chatId,
		"action":  action,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// OpenChat This method should be called if the chat is opened by the user. Many useful activities depend on the chat being opened or closed (e.g., in supergroups and channels all updates are received only for opened chats)
// @param chatId Chat identifier
func (client *Client) OpenChat(chatId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "openChat",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// CloseChat This method should be called if the chat is closed by the user. Many useful activities depend on the chat being opened or closed
// @param chatId Chat identifier
func (client *Client) CloseChat(chatId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "closeChat",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ViewMessages This method should be called if messages are being viewed by the user. Many useful activities depend on whether the messages are currently being viewed or not (e.g., marking messages as read, incrementing a view counter, updating a view counter, removing deleted messages in supergroups and channels)
// @param chatId Chat identifier
// @param messageIds The identifiers of the messages being viewed
// @param forceRead True, if messages in closed chats should be marked as read
func (client *Client) ViewMessages(chatId int64, messageIds []int64, forceRead bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "viewMessages",
		"chat_id":     chatId,
		"message_ids": messageIds,
		"force_read":  forceRead,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// OpenMessageContent This method should be called if the message content has been opened (e.g., the user has opened a photo, video, document, location or venue, or has listened to an audio file or voice note message). An updateMessageContentOpened update will be generated if something has changed
// @param chatId Chat identifier of the message
// @param messageId Identifier of the message with the opened content
func (client *Client) OpenMessageContent(chatId int64, messageId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "openMessageContent",
		"chat_id":    chatId,
		"message_id": messageId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ReadAllChatMentions Marks all mentions in a chat as read
// @param chatId Chat identifier
func (client *Client) ReadAllChatMentions(chatId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "readAllChatMentions",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// CreatePrivateChat Returns an existing chat corresponding to a given user
// @param userId User identifier
// @param force If true, the chat will be created without network request. In this case all information about the chat except its type, title and photo can be incorrect
func (client *Client) CreatePrivateChat(userId int32, force bool) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "createPrivateChat",
		"user_id": userId,
		"force":   force,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chat Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err

}

// CreateBasicGroupChat Returns an existing chat corresponding to a known basic group
// @param basicGroupId Basic group identifier
// @param force If true, the chat will be created without network request. In this case all information about the chat except its type, title and photo can be incorrect
func (client *Client) CreateBasicGroupChat(basicGroupId int32, force bool) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "createBasicGroupChat",
		"basic_group_id": basicGroupId,
		"force":          force,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chat Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err

}

// CreateSupergroupChat Returns an existing chat corresponding to a known supergroup or channel
// @param supergroupId Supergroup or channel identifier
// @param force If true, the chat will be created without network request. In this case all information about the chat except its type, title and photo can be incorrect
func (client *Client) CreateSupergroupChat(supergroupId int32, force bool) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "createSupergroupChat",
		"supergroup_id": supergroupId,
		"force":         force,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chat Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err

}

// CreateSecretChat Returns an existing chat corresponding to a known secret chat
// @param secretChatId Secret chat identifier
func (client *Client) CreateSecretChat(secretChatId int32) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "createSecretChat",
		"secret_chat_id": secretChatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatDummy Chat
	err = json.Unmarshal(result.Raw, &chatDummy)
	return &chatDummy, err

}

// CreateNewBasicGroupChat Creates a new basic group and sends a corresponding messageBasicGroupChatCreate. Returns the newly created chat
// @param userIds Identifiers of users to be added to the basic group
// @param title Title of the new basic group; 1-255 characters
func (client *Client) CreateNewBasicGroupChat(userIds []int32, title string) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "createNewBasicGroupChat",
		"user_ids": userIds,
		"title":    title,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chat Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err

}

// CreateNewSupergroupChat Creates a new supergroup or channel and sends a corresponding messageSupergroupChatCreate. Returns the newly created chat
// @param title Title of the new chat; 1-255 characters
// @param isChannel True, if a channel chat should be created
// @param description
func (client *Client) CreateNewSupergroupChat(title string, isChannel bool, description string) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "createNewSupergroupChat",
		"title":       title,
		"is_channel":  isChannel,
		"description": description,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chat Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err

}

// CreateNewSecretChat Creates a new secret chat. Returns the newly created chat
// @param userId Identifier of the target user
func (client *Client) CreateNewSecretChat(userId int32) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "createNewSecretChat",
		"user_id": userId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chat Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err

}

// UpgradeBasicGroupChatToSupergroupChat Creates a new supergroup from an existing basic group and sends a corresponding messageChatUpgradeTo and messageChatUpgradeFrom. Deactivates the original basic group
// @param chatId Identifier of the chat to upgrade
func (client *Client) UpgradeBasicGroupChatToSupergroupChat(chatId int64) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "upgradeBasicGroupChatToSupergroupChat",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatDummy Chat
	err = json.Unmarshal(result.Raw, &chatDummy)
	return &chatDummy, err

}

// SetChatTitle Changes the chat title. Supported only for basic groups, supergroups and channels. Requires administrator rights in basic groups and the appropriate administrator rights in supergroups and channels. The title will not be changed until the request to the server has been completed
// @param chatId Chat identifier
// @param title New title of the chat; 1-255 characters
func (client *Client) SetChatTitle(chatId int64, title string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "setChatTitle",
		"chat_id": chatId,
		"title":   title,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetChatPhoto Changes the photo of a chat. Supported only for basic groups, supergroups and channels. Requires administrator rights in basic groups and the appropriate administrator rights in supergroups and channels. The photo will not be changed before request to the server has been completed
// @param chatId Chat identifier
// @param photo New chat photo. You can use a zero InputFileId to delete the chat photo. Files that are accessible only by HTTP URL are not acceptable
func (client *Client) SetChatPhoto(chatId int64, photo InputFile) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "setChatPhoto",
		"chat_id": chatId,
		"photo":   photo,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetChatDraftMessage Changes the draft message in a chat
// @param chatId Chat identifier
// @param draftMessage New draft message; may be null
func (client *Client) SetChatDraftMessage(chatId int64, draftMessage *DraftMessage) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "setChatDraftMessage",
		"chat_id":       chatId,
		"draft_message": draftMessage,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetChatNotificationSettings Changes the notification settings of a chat
// @param chatId Chat identifier
// @param notificationSettings New notification settings for the chat
func (client *Client) SetChatNotificationSettings(chatId int64, notificationSettings *ChatNotificationSettings) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "setChatNotificationSettings",
		"chat_id":               chatId,
		"notification_settings": notificationSettings,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ToggleChatIsPinned Changes the pinned state of a chat. You can pin up to GetOption("pinned_chat_count_max") non-secret chats and the same number of secret chats
// @param chatId Chat identifier
// @param isPinned New value of is_pinned
func (client *Client) ToggleChatIsPinned(chatId int64, isPinned bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "toggleChatIsPinned",
		"chat_id":   chatId,
		"is_pinned": isPinned,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ToggleChatIsMarkedAsUnread Changes the marked as unread state of a chat
// @param chatId Chat identifier
// @param isMarkedAsUnread New value of is_marked_as_unread
func (client *Client) ToggleChatIsMarkedAsUnread(chatId int64, isMarkedAsUnread bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":               "toggleChatIsMarkedAsUnread",
		"chat_id":             chatId,
		"is_marked_as_unread": isMarkedAsUnread,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ToggleChatDefaultDisableNotification Changes the value of the default disable_notification parameter, used when a message is sent to a chat
// @param chatId Chat identifier
// @param defaultDisableNotification New value of default_disable_notification
func (client *Client) ToggleChatDefaultDisableNotification(chatId int64, defaultDisableNotification bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                        "toggleChatDefaultDisableNotification",
		"chat_id":                      chatId,
		"default_disable_notification": defaultDisableNotification,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetChatClientData Changes client data associated with a chat
// @param chatId Chat identifier
// @param clientData New value of client_data
func (client *Client) SetChatClientData(chatId int64, clientData string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "setChatClientData",
		"chat_id":     chatId,
		"client_data": clientData,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// JoinChat Adds current user as a new member to a chat. Private and secret chats can't be joined using this method
// @param chatId Chat identifier
func (client *Client) JoinChat(chatId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "joinChat",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// LeaveChat Removes current user from chat members. Private and secret chats can't be left using this method
// @param chatId Chat identifier
func (client *Client) LeaveChat(chatId int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "leaveChat",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// AddChatMember Adds a new member to a chat. Members can't be added to private or secret chats. Members will not be added until the chat state has been synchronized with the server
// @param chatId Chat identifier
// @param userId Identifier of the user
// @param forwardLimit The number of earlier messages from the chat to be forwarded to the new member; up to 300. Ignored for supergroups and channels
func (client *Client) AddChatMember(chatId int64, userId int32, forwardLimit int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "addChatMember",
		"chat_id":       chatId,
		"user_id":       userId,
		"forward_limit": forwardLimit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// AddChatMembers Adds multiple new members to a chat. Currently this option is only available for supergroups and channels. This option can't be used to join a chat. Members can't be added to a channel if it has more than 200 members. Members will not be added until the chat state has been synchronized with the server
// @param chatId Chat identifier
// @param userIds Identifiers of the users to be added to the chat
func (client *Client) AddChatMembers(chatId int64, userIds []int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "addChatMembers",
		"chat_id":  chatId,
		"user_ids": userIds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetChatMemberStatus Changes the status of a chat member, needs appropriate privileges. This function is currently not suitable for adding new members to the chat; instead, use addChatMember. The chat member status will not be changed until it has been synchronized with the server
// @param chatId Chat identifier
// @param userId User identifier
// @param status The new status of the member in the chat
func (client *Client) SetChatMemberStatus(chatId int64, userId int32, status ChatMemberStatus) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "setChatMemberStatus",
		"chat_id": chatId,
		"user_id": userId,
		"status":  status,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetChatMember Returns information about a single member of a chat
// @param chatId Chat identifier
// @param userId User identifier
func (client *Client) GetChatMember(chatId int64, userId int32) (*ChatMember, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getChatMember",
		"chat_id": chatId,
		"user_id": userId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatMember ChatMember
	err = json.Unmarshal(result.Raw, &chatMember)
	return &chatMember, err

}

// SearchChatMembers Searches for a specified query in the first name, last name and username of the members of a specified chat. Requires administrator rights in channels
// @param chatId Chat identifier
// @param query Query to search for
// @param limit The maximum number of users to be returned
// @param filter The type of users to return. By default, chatMembersFilterMembers
func (client *Client) SearchChatMembers(chatId int64, query string, limit int32, filter ChatMembersFilter) (*ChatMembers, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "searchChatMembers",
		"chat_id": chatId,
		"query":   query,
		"limit":   limit,
		"filter":  filter,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatMembers ChatMembers
	err = json.Unmarshal(result.Raw, &chatMembers)
	return &chatMembers, err

}

// GetChatAdministrators Returns a list of users who are administrators of the chat
// @param chatId Chat identifier
func (client *Client) GetChatAdministrators(chatId int64) (*Users, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getChatAdministrators",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var users Users
	err = json.Unmarshal(result.Raw, &users)
	return &users, err

}

// ClearAllDraftMessages Clears draft messages in all chats
// @param excludeSecretChats If true, local draft messages in secret chats will not be cleared
func (client *Client) ClearAllDraftMessages(excludeSecretChats bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                "clearAllDraftMessages",
		"exclude_secret_chats": excludeSecretChats,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetScopeNotificationSettings Returns the notification settings for chats of a given type
// @param scope Types of chats for which to return the notification settings information
func (client *Client) GetScopeNotificationSettings(scope NotificationSettingsScope) (*ScopeNotificationSettings, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getScopeNotificationSettings",
		"scope": scope,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var scopeNotificationSettings ScopeNotificationSettings
	err = json.Unmarshal(result.Raw, &scopeNotificationSettings)
	return &scopeNotificationSettings, err

}

// SetScopeNotificationSettings Changes notification settings for chats of a given type
// @param scope Types of chats for which to change the notification settings
// @param notificationSettings The new notification settings for the given scope
func (client *Client) SetScopeNotificationSettings(scope NotificationSettingsScope, notificationSettings *ScopeNotificationSettings) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "setScopeNotificationSettings",
		"scope":                 scope,
		"notification_settings": notificationSettings,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ResetAllNotificationSettings Resets all notification settings to their default values. By default, all chats are unmuted, the sound is set to "default" and message previews are shown
func (client *Client) ResetAllNotificationSettings() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "resetAllNotificationSettings",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetPinnedChats Changes the order of pinned chats
// @param chatIds The new list of pinned chats
func (client *Client) SetPinnedChats(chatIds []int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "setPinnedChats",
		"chat_ids": chatIds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// DownloadFile Asynchronously downloads a file from the cloud. updateFile will be used to notify about the download progress and successful completion of the download. Returns file state just after the download has been started
// @param fileId Identifier of the file to download
// @param priority Priority of the download (1-32). The higher the priority, the earlier the file will be downloaded. If the priorities of two files are equal, then the last one for which downloadFile was called will be downloaded first
func (client *Client) DownloadFile(fileId int32, priority int32) (*File, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "downloadFile",
		"file_id":  fileId,
		"priority": priority,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var fileDummy File
	err = json.Unmarshal(result.Raw, &fileDummy)
	return &fileDummy, err

}

// CancelDownloadFile Stops the downloading of a file. If a file has already been downloaded, does nothing
// @param fileId Identifier of a file to stop downloading
// @param onlyIfPending Pass true to stop downloading only if it hasn't been started, i.e. request hasn't been sent to server
func (client *Client) CancelDownloadFile(fileId int32, onlyIfPending bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "cancelDownloadFile",
		"file_id":         fileId,
		"only_if_pending": onlyIfPending,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// UploadFile Asynchronously uploads a file to the cloud without sending it in a message. updateFile will be used to notify about upload progress and successful completion of the upload. The file will not have a persistent remote identifier until it will be sent in a message
// @param file File to upload
// @param fileType File type
// @param priority Priority of the upload (1-32). The higher the priority, the earlier the file will be uploaded. If the priorities of two files are equal, then the first one for which uploadFile was called will be uploaded first
func (client *Client) UploadFile(file InputFile, fileType FileType, priority int32) (*File, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "uploadFile",
		"file":      file,
		"file_type": fileType,
		"priority":  priority,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var fileDummy File
	err = json.Unmarshal(result.Raw, &fileDummy)
	return &fileDummy, err

}

// CancelUploadFile Stops the uploading of a file. Supported only for files uploaded by using uploadFile. For other files the behavior is undefined
// @param fileId Identifier of the file to stop uploading
func (client *Client) CancelUploadFile(fileId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "cancelUploadFile",
		"file_id": fileId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetFileGenerationProgress The next part of a file was generated
// @param generationId The identifier of the generation process
// @param expectedSize Expected size of the generated file, in bytes; 0 if unknown
// @param localPrefixSize The number of bytes already generated
func (client *Client) SetFileGenerationProgress(generationId JSONInt64, expectedSize int32, localPrefixSize int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "setFileGenerationProgress",
		"generation_id":     generationId,
		"expected_size":     expectedSize,
		"local_prefix_size": localPrefixSize,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// FinishFileGeneration Finishes the file generation
// @param generationId The identifier of the generation process
// @param error If set, means that file generation has failed and should be terminated
func (client *Client) FinishFileGeneration(generationId JSONInt64, error *Error) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "finishFileGeneration",
		"generation_id": generationId,
		"error":         error,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// DeleteFile Deletes a file from the TDLib file cache
// @param fileId Identifier of the file to delete
func (client *Client) DeleteFile(fileId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "deleteFile",
		"file_id": fileId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GenerateChatInviteLink Generates a new invite link for a chat; the previously generated link is revoked. Available for basic groups, supergroups, and channels. In basic groups this can be called only by the group's creator; in supergroups and channels this requires appropriate administrator rights
// @param chatId Chat identifier
func (client *Client) GenerateChatInviteLink(chatId int64) (*ChatInviteLink, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "generateChatInviteLink",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatInviteLink ChatInviteLink
	err = json.Unmarshal(result.Raw, &chatInviteLink)
	return &chatInviteLink, err

}

// CheckChatInviteLink Checks the validity of an invite link for a chat and returns information about the corresponding chat
// @param inviteLink Invite link to be checked; should begin with "https://t.me/joinchat/", "https://telegram.me/joinchat/", or "https://telegram.dog/joinchat/"
func (client *Client) CheckChatInviteLink(inviteLink string) (*ChatInviteLinkInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "checkChatInviteLink",
		"invite_link": inviteLink,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatInviteLinkInfo ChatInviteLinkInfo
	err = json.Unmarshal(result.Raw, &chatInviteLinkInfo)
	return &chatInviteLinkInfo, err

}

// JoinChatByInviteLink Uses an invite link to add the current user to the chat if possible. The new member will not be added until the chat state has been synchronized with the server
// @param inviteLink Invite link to import; should begin with "https://t.me/joinchat/", "https://telegram.me/joinchat/", or "https://telegram.dog/joinchat/"
func (client *Client) JoinChatByInviteLink(inviteLink string) (*Chat, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "joinChatByInviteLink",
		"invite_link": inviteLink,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chat Chat
	err = json.Unmarshal(result.Raw, &chat)
	return &chat, err

}

// CreateCall Creates a new call
// @param userId Identifier of the user to be called
// @param protocol Description of the call protocols supported by the client
func (client *Client) CreateCall(userId int32, protocol *CallProtocol) (*CallId, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "createCall",
		"user_id":  userId,
		"protocol": protocol,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var callId CallId
	err = json.Unmarshal(result.Raw, &callId)
	return &callId, err

}

// AcceptCall Accepts an incoming call
// @param callId Call identifier
// @param protocol Description of the call protocols supported by the client
func (client *Client) AcceptCall(callId int32, protocol *CallProtocol) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "acceptCall",
		"call_id":  callId,
		"protocol": protocol,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// DiscardCall Discards a call
// @param callId Call identifier
// @param isDisconnected True, if the user was disconnected
// @param duration The call duration, in seconds
// @param connectionId Identifier of the connection used during the call
func (client *Client) DiscardCall(callId int32, isDisconnected bool, duration int32, connectionId JSONInt64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "discardCall",
		"call_id":         callId,
		"is_disconnected": isDisconnected,
		"duration":        duration,
		"connection_id":   connectionId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SendCallRating Sends a call rating
// @param callId Call identifier
// @param rating Call rating; 1-5
// @param comment An optional user comment if the rating is less than 5
func (client *Client) SendCallRating(callId int32, rating int32, comment string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "sendCallRating",
		"call_id": callId,
		"rating":  rating,
		"comment": comment,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SendCallDebugInformation Sends debug information for a call
// @param callId Call identifier
// @param debugInformation Debug information in application-specific format
func (client *Client) SendCallDebugInformation(callId int32, debugInformation string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "sendCallDebugInformation",
		"call_id":           callId,
		"debug_information": debugInformation,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// BlockUser Adds a user to the blacklist
// @param userId User identifier
func (client *Client) BlockUser(userId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "blockUser",
		"user_id": userId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// UnblockUser Removes a user from the blacklist
// @param userId User identifier
func (client *Client) UnblockUser(userId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "unblockUser",
		"user_id": userId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetBlockedUsers Returns users that were blocked by the current user
// @param offset Number of users to skip in the result; must be non-negative
// @param limit Maximum number of users to return; up to 100
func (client *Client) GetBlockedUsers(offset int32, limit int32) (*Users, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":  "getBlockedUsers",
		"offset": offset,
		"limit":  limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var users Users
	err = json.Unmarshal(result.Raw, &users)
	return &users, err

}

// ImportContacts Adds new contacts or edits existing contacts; contacts' user identifiers are ignored
// @param contacts The list of contacts to import or edit, contact's vCard are ignored and are not imported
func (client *Client) ImportContacts(contacts []Contact) (*ImportedContacts, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "importContacts",
		"contacts": contacts,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var importedContacts ImportedContacts
	err = json.Unmarshal(result.Raw, &importedContacts)
	return &importedContacts, err

}

// GetContacts Returns all user contacts
func (client *Client) GetContacts() (*Users, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getContacts",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var users Users
	err = json.Unmarshal(result.Raw, &users)
	return &users, err

}

// SearchContacts Searches for the specified query in the first names, last names and usernames of the known user contacts
// @param query Query to search for; can be empty to return all contacts
// @param limit Maximum number of users to be returned
func (client *Client) SearchContacts(query string, limit int32) (*Users, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "searchContacts",
		"query": query,
		"limit": limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var users Users
	err = json.Unmarshal(result.Raw, &users)
	return &users, err

}

// RemoveContacts Removes users from the contacts list
// @param userIds Identifiers of users to be deleted
func (client *Client) RemoveContacts(userIds []int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "removeContacts",
		"user_ids": userIds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetImportedContactCount Returns the total number of imported contacts
func (client *Client) GetImportedContactCount() (*Count, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getImportedContactCount",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var count Count
	err = json.Unmarshal(result.Raw, &count)
	return &count, err

}

// ChangeImportedContacts Changes imported contacts using the list of current user contacts saved on the device. Imports newly added contacts and, if at least the file database is enabled, deletes recently deleted contacts.
// @param contacts
func (client *Client) ChangeImportedContacts(contacts []Contact) (*ImportedContacts, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "changeImportedContacts",
		"contacts": contacts,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var importedContacts ImportedContacts
	err = json.Unmarshal(result.Raw, &importedContacts)
	return &importedContacts, err

}

// ClearImportedContacts Clears all imported contacts, contacts list remains unchanged
func (client *Client) ClearImportedContacts() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "clearImportedContacts",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetUserProfilePhotos Returns the profile photos of a user. The result of this query may be outdated: some photos might have been deleted already
// @param userId User identifier
// @param offset The number of photos to skip; must be non-negative
// @param limit Maximum number of photos to be returned; up to 100
func (client *Client) GetUserProfilePhotos(userId int32, offset int32, limit int32) (*UserProfilePhotos, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getUserProfilePhotos",
		"user_id": userId,
		"offset":  offset,
		"limit":   limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var userProfilePhotos UserProfilePhotos
	err = json.Unmarshal(result.Raw, &userProfilePhotos)
	return &userProfilePhotos, err

}

// GetStickers Returns stickers from the installed sticker sets that correspond to a given emoji. If the emoji is not empty, favorite and recently used stickers may also be returned
// @param emoji String representation of emoji. If empty, returns all known installed stickers
// @param limit Maximum number of stickers to be returned
func (client *Client) GetStickers(emoji string, limit int32) (*Stickers, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getStickers",
		"emoji": emoji,
		"limit": limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickers Stickers
	err = json.Unmarshal(result.Raw, &stickers)
	return &stickers, err

}

// SearchStickers Searches for stickers from public sticker sets that correspond to a given emoji
// @param emoji String representation of emoji; must be non-empty
// @param limit Maximum number of stickers to be returned
func (client *Client) SearchStickers(emoji string, limit int32) (*Stickers, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "searchStickers",
		"emoji": emoji,
		"limit": limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickers Stickers
	err = json.Unmarshal(result.Raw, &stickers)
	return &stickers, err

}

// GetInstalledStickerSets Returns a list of installed sticker sets
// @param isMasks Pass true to return mask sticker sets; pass false to return ordinary sticker sets
func (client *Client) GetInstalledStickerSets(isMasks bool) (*StickerSets, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "getInstalledStickerSets",
		"is_masks": isMasks,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSets StickerSets
	err = json.Unmarshal(result.Raw, &stickerSets)
	return &stickerSets, err

}

// GetArchivedStickerSets Returns a list of archived sticker sets
// @param isMasks Pass true to return mask stickers sets; pass false to return ordinary sticker sets
// @param offsetStickerSetId Identifier of the sticker set from which to return the result
// @param limit Maximum number of sticker sets to return
func (client *Client) GetArchivedStickerSets(isMasks bool, offsetStickerSetId JSONInt64, limit int32) (*StickerSets, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                 "getArchivedStickerSets",
		"is_masks":              isMasks,
		"offset_sticker_set_id": offsetStickerSetId,
		"limit":                 limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSets StickerSets
	err = json.Unmarshal(result.Raw, &stickerSets)
	return &stickerSets, err

}

// GetTrendingStickerSets Returns a list of trending sticker sets
func (client *Client) GetTrendingStickerSets() (*StickerSets, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getTrendingStickerSets",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSets StickerSets
	err = json.Unmarshal(result.Raw, &stickerSets)
	return &stickerSets, err

}

// GetAttachedStickerSets Returns a list of sticker sets attached to a file. Currently only photos and videos can have attached sticker sets
// @param fileId File identifier
func (client *Client) GetAttachedStickerSets(fileId int32) (*StickerSets, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getAttachedStickerSets",
		"file_id": fileId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSets StickerSets
	err = json.Unmarshal(result.Raw, &stickerSets)
	return &stickerSets, err

}

// GetStickerSet Returns information about a sticker set by its identifier
// @param setId Identifier of the sticker set
func (client *Client) GetStickerSet(setId JSONInt64) (*StickerSet, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":  "getStickerSet",
		"set_id": setId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSet StickerSet
	err = json.Unmarshal(result.Raw, &stickerSet)
	return &stickerSet, err

}

// SearchStickerSet Searches for a sticker set by its name
// @param name Name of the sticker set
func (client *Client) SearchStickerSet(name string) (*StickerSet, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "searchStickerSet",
		"name":  name,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSet StickerSet
	err = json.Unmarshal(result.Raw, &stickerSet)
	return &stickerSet, err

}

// SearchInstalledStickerSets Searches for installed sticker sets by looking for specified query in their title and name
// @param isMasks Pass true to return mask sticker sets; pass false to return ordinary sticker sets
// @param query Query to search for
// @param limit Maximum number of sticker sets to return
func (client *Client) SearchInstalledStickerSets(isMasks bool, query string, limit int32) (*StickerSets, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "searchInstalledStickerSets",
		"is_masks": isMasks,
		"query":    query,
		"limit":    limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSets StickerSets
	err = json.Unmarshal(result.Raw, &stickerSets)
	return &stickerSets, err

}

// SearchStickerSets Searches for ordinary sticker sets by looking for specified query in their title and name. Excludes installed sticker sets from the results
// @param query Query to search for
func (client *Client) SearchStickerSets(query string) (*StickerSets, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "searchStickerSets",
		"query": query,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSets StickerSets
	err = json.Unmarshal(result.Raw, &stickerSets)
	return &stickerSets, err

}

// ChangeStickerSet Installs/uninstalls or activates/archives a sticker set
// @param setId Identifier of the sticker set
// @param isInstalled The new value of is_installed
// @param isArchived The new value of is_archived. A sticker set can't be installed and archived simultaneously
func (client *Client) ChangeStickerSet(setId JSONInt64, isInstalled bool, isArchived bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "changeStickerSet",
		"set_id":       setId,
		"is_installed": isInstalled,
		"is_archived":  isArchived,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ViewTrendingStickerSets Informs the server that some trending sticker sets have been viewed by the user
// @param stickerSetIds Identifiers of viewed trending sticker sets
func (client *Client) ViewTrendingStickerSets(stickerSetIds []JSONInt64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "viewTrendingStickerSets",
		"sticker_set_ids": stickerSetIds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ReorderInstalledStickerSets Changes the order of installed sticker sets
// @param isMasks Pass true to change the order of mask sticker sets; pass false to change the order of ordinary sticker sets
// @param stickerSetIds Identifiers of installed sticker sets in the new correct order
func (client *Client) ReorderInstalledStickerSets(isMasks bool, stickerSetIds []JSONInt64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "reorderInstalledStickerSets",
		"is_masks":        isMasks,
		"sticker_set_ids": stickerSetIds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetRecentStickers Returns a list of recently used stickers
// @param isAttached Pass true to return stickers and masks that were recently attached to photos or video files; pass false to return recently sent stickers
func (client *Client) GetRecentStickers(isAttached bool) (*Stickers, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "getRecentStickers",
		"is_attached": isAttached,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickers Stickers
	err = json.Unmarshal(result.Raw, &stickers)
	return &stickers, err

}

// AddRecentSticker Manually adds a new sticker to the list of recently used stickers. The new sticker is added to the top of the list. If the sticker was already in the list, it is removed from the list first. Only stickers belonging to a sticker set can be added to this list
// @param isAttached Pass true to add the sticker to the list of stickers recently attached to photo or video files; pass false to add the sticker to the list of recently sent stickers
// @param sticker Sticker file to add
func (client *Client) AddRecentSticker(isAttached bool, sticker InputFile) (*Stickers, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "addRecentSticker",
		"is_attached": isAttached,
		"sticker":     sticker,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickers Stickers
	err = json.Unmarshal(result.Raw, &stickers)
	return &stickers, err

}

// RemoveRecentSticker Removes a sticker from the list of recently used stickers
// @param isAttached Pass true to remove the sticker from the list of stickers recently attached to photo or video files; pass false to remove the sticker from the list of recently sent stickers
// @param sticker Sticker file to delete
func (client *Client) RemoveRecentSticker(isAttached bool, sticker InputFile) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "removeRecentSticker",
		"is_attached": isAttached,
		"sticker":     sticker,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ClearRecentStickers Clears the list of recently used stickers
// @param isAttached Pass true to clear the list of stickers recently attached to photo or video files; pass false to clear the list of recently sent stickers
func (client *Client) ClearRecentStickers(isAttached bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "clearRecentStickers",
		"is_attached": isAttached,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetFavoriteStickers Returns favorite stickers
func (client *Client) GetFavoriteStickers() (*Stickers, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getFavoriteStickers",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickers Stickers
	err = json.Unmarshal(result.Raw, &stickers)
	return &stickers, err

}

// AddFavoriteSticker Adds a new sticker to the list of favorite stickers. The new sticker is added to the top of the list. If the sticker was already in the list, it is removed from the list first. Only stickers belonging to a sticker set can be added to this list
// @param sticker Sticker file to add
func (client *Client) AddFavoriteSticker(sticker InputFile) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "addFavoriteSticker",
		"sticker": sticker,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// RemoveFavoriteSticker Removes a sticker from the list of favorite stickers
// @param sticker Sticker file to delete from the list
func (client *Client) RemoveFavoriteSticker(sticker InputFile) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "removeFavoriteSticker",
		"sticker": sticker,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetStickerEmojis Returns emoji corresponding to a sticker
// @param sticker Sticker file identifier
func (client *Client) GetStickerEmojis(sticker InputFile) (*StickerEmojis, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getStickerEmojis",
		"sticker": sticker,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerEmojis StickerEmojis
	err = json.Unmarshal(result.Raw, &stickerEmojis)
	return &stickerEmojis, err

}

// GetSavedAnimations Returns saved animations
func (client *Client) GetSavedAnimations() (*Animations, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getSavedAnimations",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var animations Animations
	err = json.Unmarshal(result.Raw, &animations)
	return &animations, err

}

// AddSavedAnimation Manually adds a new animation to the list of saved animations. The new animation is added to the beginning of the list. If the animation was already in the list, it is removed first. Only non-secret video animations with MIME type "video/mp4" can be added to the list
// @param animation The animation file to be added. Only animations known to the server (i.e. successfully sent via a message) can be added to the list
func (client *Client) AddSavedAnimation(animation InputFile) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "addSavedAnimation",
		"animation": animation,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// RemoveSavedAnimation Removes an animation from the list of saved animations
// @param animation Animation file to be removed
func (client *Client) RemoveSavedAnimation(animation InputFile) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":     "removeSavedAnimation",
		"animation": animation,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetRecentInlineBots Returns up to 20 recently used inline bots in the order of their last usage
func (client *Client) GetRecentInlineBots() (*Users, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getRecentInlineBots",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var users Users
	err = json.Unmarshal(result.Raw, &users)
	return &users, err

}

// SearchHashtags Searches for recently used hashtags by their prefix
// @param prefix Hashtag prefix to search for
// @param limit Maximum number of hashtags to be returned
func (client *Client) SearchHashtags(prefix string, limit int32) (*Hashtags, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":  "searchHashtags",
		"prefix": prefix,
		"limit":  limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var hashtags Hashtags
	err = json.Unmarshal(result.Raw, &hashtags)
	return &hashtags, err

}

// RemoveRecentHashtag Removes a hashtag from the list of recently used hashtags
// @param hashtag Hashtag to delete
func (client *Client) RemoveRecentHashtag(hashtag string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "removeRecentHashtag",
		"hashtag": hashtag,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetWebPagePreview Returns a web page preview by the text of the message. Do not call this function too often. Returns a 404 error if the web page has no preview
// @param text Message text with formatting
func (client *Client) GetWebPagePreview(text *FormattedText) (*WebPage, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getWebPagePreview",
		"text":  text,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var webPage WebPage
	err = json.Unmarshal(result.Raw, &webPage)
	return &webPage, err

}

// GetWebPageInstantView Returns an instant view version of a web page if available. Returns a 404 error if the web page has no instant view page
// @param url The web page URL
// @param forceFull If true, the full instant view for the web page will be returned
func (client *Client) GetWebPageInstantView(url string, forceFull bool) (*WebPageInstantView, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getWebPageInstantView",
		"url":        url,
		"force_full": forceFull,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var webPageInstantView WebPageInstantView
	err = json.Unmarshal(result.Raw, &webPageInstantView)
	return &webPageInstantView, err

}

// SetProfilePhoto Uploads a new profile photo for the current user. If something changes, updateUser will be sent
// @param photo Profile photo to set. inputFileId and inputFileRemote may still be unsupported
func (client *Client) SetProfilePhoto(photo InputFile) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "setProfilePhoto",
		"photo": photo,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// DeleteProfilePhoto Deletes a profile photo. If something changes, updateUser will be sent
// @param profilePhotoId Identifier of the profile photo to delete
func (client *Client) DeleteProfilePhoto(profilePhotoId JSONInt64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":            "deleteProfilePhoto",
		"profile_photo_id": profilePhotoId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetName Changes the first and last name of the current user. If something changes, updateUser will be sent
// @param firstName The new value of the first name for the user; 1-255 characters
// @param lastName The new value of the optional last name for the user; 0-255 characters
func (client *Client) SetName(firstName string, lastName string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "setName",
		"first_name": firstName,
		"last_name":  lastName,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetBio Changes the bio of the current user
// @param bio The new value of the user bio; 0-70 characters without line feeds
func (client *Client) SetBio(bio string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "setBio",
		"bio":   bio,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetUsername Changes the username of the current user. If something changes, updateUser will be sent
// @param username The new value of the username. Use an empty string to remove the username
func (client *Client) SetUsername(username string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "setUsername",
		"username": username,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ChangePhoneNumber Changes the phone number of the user and sends an authentication code to the user's new phone number. On success, returns information about the sent code
// @param phoneNumber The new phone number of the user in international format
// @param allowFlashCall Pass true if the code can be sent via flash call to the specified phone number
// @param isCurrentPhoneNumber Pass true if the phone number is used on the current device. Ignored if allow_flash_call is false
func (client *Client) ChangePhoneNumber(phoneNumber string, allowFlashCall bool, isCurrentPhoneNumber bool) (*AuthenticationCodeInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                   "changePhoneNumber",
		"phone_number":            phoneNumber,
		"allow_flash_call":        allowFlashCall,
		"is_current_phone_number": isCurrentPhoneNumber,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var authenticationCodeInfo AuthenticationCodeInfo
	err = json.Unmarshal(result.Raw, &authenticationCodeInfo)
	return &authenticationCodeInfo, err

}

// ResendChangePhoneNumberCode Re-sends the authentication code sent to confirm a new phone number for the user. Works only if the previously received authenticationCodeInfo next_code_type was not null
func (client *Client) ResendChangePhoneNumberCode() (*AuthenticationCodeInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "resendChangePhoneNumberCode",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var authenticationCodeInfo AuthenticationCodeInfo
	err = json.Unmarshal(result.Raw, &authenticationCodeInfo)
	return &authenticationCodeInfo, err

}

// CheckChangePhoneNumberCode Checks the authentication code sent to confirm a new phone number of the user
// @param code Verification code received by SMS, phone call or flash call
func (client *Client) CheckChangePhoneNumberCode(code string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "checkChangePhoneNumberCode",
		"code":  code,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetActiveSessions Returns all active sessions of the current user
func (client *Client) GetActiveSessions() (*Sessions, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getActiveSessions",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var sessions Sessions
	err = json.Unmarshal(result.Raw, &sessions)
	return &sessions, err

}

// TerminateSession Terminates a session of the current user
// @param sessionId Session identifier
func (client *Client) TerminateSession(sessionId JSONInt64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "terminateSession",
		"session_id": sessionId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// TerminateAllOtherSessions Terminates all other sessions of the current user
func (client *Client) TerminateAllOtherSessions() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "terminateAllOtherSessions",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetConnectedWebsites Returns all website where the current user used Telegram to log in
func (client *Client) GetConnectedWebsites() (*ConnectedWebsites, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getConnectedWebsites",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var connectedWebsites ConnectedWebsites
	err = json.Unmarshal(result.Raw, &connectedWebsites)
	return &connectedWebsites, err

}

// DisconnectWebsite Disconnects website from the current user's Telegram account
// @param websiteId Website identifier
func (client *Client) DisconnectWebsite(websiteId JSONInt64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "disconnectWebsite",
		"website_id": websiteId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// DisconnectAllWebsites Disconnects all websites from the current user's Telegram account
func (client *Client) DisconnectAllWebsites() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "disconnectAllWebsites",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ToggleBasicGroupAdministrators Toggles the "All members are admins" setting in basic groups; requires creator privileges in the group
// @param basicGroupId Identifier of the basic group
// @param everyoneIsAdministrator New value of everyone_is_administrator
func (client *Client) ToggleBasicGroupAdministrators(basicGroupId int32, everyoneIsAdministrator bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                     "toggleBasicGroupAdministrators",
		"basic_group_id":            basicGroupId,
		"everyone_is_administrator": everyoneIsAdministrator,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetSupergroupUsername Changes the username of a supergroup or channel, requires creator privileges in the supergroup or channel
// @param supergroupId Identifier of the supergroup or channel
// @param username New value of the username. Use an empty string to remove the username
func (client *Client) SetSupergroupUsername(supergroupId int32, username string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "setSupergroupUsername",
		"supergroup_id": supergroupId,
		"username":      username,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetSupergroupStickerSet Changes the sticker set of a supergroup; requires appropriate rights in the supergroup
// @param supergroupId Identifier of the supergroup
// @param stickerSetId New value of the supergroup sticker set identifier. Use 0 to remove the supergroup sticker set
func (client *Client) SetSupergroupStickerSet(supergroupId int32, stickerSetId JSONInt64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "setSupergroupStickerSet",
		"supergroup_id":  supergroupId,
		"sticker_set_id": stickerSetId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ToggleSupergroupInvites Toggles whether all members of a supergroup can add new members; requires appropriate administrator rights in the supergroup.
// @param supergroupId Identifier of the supergroup
// @param anyoneCanInvite New value of anyone_can_invite
func (client *Client) ToggleSupergroupInvites(supergroupId int32, anyoneCanInvite bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":             "toggleSupergroupInvites",
		"supergroup_id":     supergroupId,
		"anyone_can_invite": anyoneCanInvite,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ToggleSupergroupSignMessages Toggles sender signatures messages sent in a channel; requires appropriate administrator rights in the channel.
// @param supergroupId Identifier of the channel
// @param signMessages New value of sign_messages
func (client *Client) ToggleSupergroupSignMessages(supergroupId int32, signMessages bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "toggleSupergroupSignMessages",
		"supergroup_id": supergroupId,
		"sign_messages": signMessages,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ToggleSupergroupIsAllHistoryAvailable Toggles whether the message history of a supergroup is available to new members; requires appropriate administrator rights in the supergroup.
// @param supergroupId The identifier of the supergroup
// @param isAllHistoryAvailable The new value of is_all_history_available
func (client *Client) ToggleSupergroupIsAllHistoryAvailable(supergroupId int32, isAllHistoryAvailable bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                    "toggleSupergroupIsAllHistoryAvailable",
		"supergroup_id":            supergroupId,
		"is_all_history_available": isAllHistoryAvailable,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetSupergroupDescription Changes information about a supergroup or channel; requires appropriate administrator rights
// @param supergroupId Identifier of the supergroup or channel
// @param description
func (client *Client) SetSupergroupDescription(supergroupId int32, description string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "setSupergroupDescription",
		"supergroup_id": supergroupId,
		"description":   description,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// PinSupergroupMessage Pins a message in a supergroup or channel; requires appropriate administrator rights in the supergroup or channel
// @param supergroupId Identifier of the supergroup or channel
// @param messageId Identifier of the new pinned message
// @param disableNotification True, if there should be no notification about the pinned message
func (client *Client) PinSupergroupMessage(supergroupId int32, messageId int64, disableNotification bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                "pinSupergroupMessage",
		"supergroup_id":        supergroupId,
		"message_id":           messageId,
		"disable_notification": disableNotification,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// UnpinSupergroupMessage Removes the pinned message from a supergroup or channel; requires appropriate administrator rights in the supergroup or channel
// @param supergroupId Identifier of the supergroup or channel
func (client *Client) UnpinSupergroupMessage(supergroupId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "unpinSupergroupMessage",
		"supergroup_id": supergroupId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ReportSupergroupSpam Reports some messages from a user in a supergroup as spam; requires administrator rights in the supergroup
// @param supergroupId Supergroup identifier
// @param userId User identifier
// @param messageIds Identifiers of messages sent in the supergroup by the user. This list must be non-empty
func (client *Client) ReportSupergroupSpam(supergroupId int32, userId int32, messageIds []int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "reportSupergroupSpam",
		"supergroup_id": supergroupId,
		"user_id":       userId,
		"message_ids":   messageIds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetSupergroupMembers Returns information about members or banned users in a supergroup or channel. Can be used only if SupergroupFullInfo.can_get_members == true; additionally, administrator privileges may be required for some filters
// @param supergroupId Identifier of the supergroup or channel
// @param filter The type of users to return. By default, supergroupMembersRecent
// @param offset Number of users to skip
// @param limit The maximum number of users be returned; up to 200
func (client *Client) GetSupergroupMembers(supergroupId int32, filter SupergroupMembersFilter, offset int32, limit int32) (*ChatMembers, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "getSupergroupMembers",
		"supergroup_id": supergroupId,
		"filter":        filter,
		"offset":        offset,
		"limit":         limit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatMembers ChatMembers
	err = json.Unmarshal(result.Raw, &chatMembers)
	return &chatMembers, err

}

// DeleteSupergroup Deletes a supergroup or channel along with all messages in the corresponding chat. This will release the supergroup or channel username and remove all members; requires creator privileges in the supergroup or channel. Chats with more than 1000 members can't be deleted using this method
// @param supergroupId Identifier of the supergroup or channel
func (client *Client) DeleteSupergroup(supergroupId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "deleteSupergroup",
		"supergroup_id": supergroupId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// CloseSecretChat Closes a secret chat, effectively transfering its state to secretChatStateClosed
// @param secretChatId Secret chat identifier
func (client *Client) CloseSecretChat(secretChatId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "closeSecretChat",
		"secret_chat_id": secretChatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetChatEventLog Returns a list of service actions taken by chat members and administrators in the last 48 hours. Available only in supergroups and channels. Requires administrator rights. Returns results in reverse chronological order (i. e., in order of decreasing event_id)
// @param chatId Chat identifier
// @param query Search query by which to filter events
// @param fromEventId Identifier of an event from which to return results. Use 0 to get results from the latest events
// @param limit Maximum number of events to return; up to 100
// @param filters The types of events to return. By default, all types will be returned
// @param userIds User identifiers by which to filter events. By default, events relating to all users will be returned
func (client *Client) GetChatEventLog(chatId int64, query string, fromEventId JSONInt64, limit int32, filters *ChatEventLogFilters, userIds []int32) (*ChatEvents, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "getChatEventLog",
		"chat_id":       chatId,
		"query":         query,
		"from_event_id": fromEventId,
		"limit":         limit,
		"filters":       filters,
		"user_ids":      userIds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatEvents ChatEvents
	err = json.Unmarshal(result.Raw, &chatEvents)
	return &chatEvents, err

}

// GetPaymentForm Returns an invoice payment form. This method should be called when the user presses inlineKeyboardButtonBuy
// @param chatId Chat identifier of the Invoice message
// @param messageId Message identifier
func (client *Client) GetPaymentForm(chatId int64, messageId int64) (*PaymentForm, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getPaymentForm",
		"chat_id":    chatId,
		"message_id": messageId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var paymentForm PaymentForm
	err = json.Unmarshal(result.Raw, &paymentForm)
	return &paymentForm, err

}

// ValidateOrderInfo Validates the order information provided by a user and returns the available shipping options for a flexible invoice
// @param chatId Chat identifier of the Invoice message
// @param messageId Message identifier
// @param orderInfo The order information, provided by the user
// @param allowSave True, if the order information can be saved
func (client *Client) ValidateOrderInfo(chatId int64, messageId int64, orderInfo *OrderInfo, allowSave bool) (*ValidatedOrderInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "validateOrderInfo",
		"chat_id":    chatId,
		"message_id": messageId,
		"order_info": orderInfo,
		"allow_save": allowSave,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var validatedOrderInfo ValidatedOrderInfo
	err = json.Unmarshal(result.Raw, &validatedOrderInfo)
	return &validatedOrderInfo, err

}

// SendPaymentForm Sends a filled-out payment form to the bot for final verification
// @param chatId Chat identifier of the Invoice message
// @param messageId Message identifier
// @param orderInfoId Identifier returned by ValidateOrderInfo, or an empty string
// @param shippingOptionId Identifier of a chosen shipping option, if applicable
// @param credentials The credentials chosen by user for payment
func (client *Client) SendPaymentForm(chatId int64, messageId int64, orderInfoId string, shippingOptionId string, credentials InputCredentials) (*PaymentResult, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":              "sendPaymentForm",
		"chat_id":            chatId,
		"message_id":         messageId,
		"order_info_id":      orderInfoId,
		"shipping_option_id": shippingOptionId,
		"credentials":        credentials,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var paymentResult PaymentResult
	err = json.Unmarshal(result.Raw, &paymentResult)
	return &paymentResult, err

}

// GetPaymentReceipt Returns information about a successful payment
// @param chatId Chat identifier of the PaymentSuccessful message
// @param messageId Message identifier
func (client *Client) GetPaymentReceipt(chatId int64, messageId int64) (*PaymentReceipt, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getPaymentReceipt",
		"chat_id":    chatId,
		"message_id": messageId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var paymentReceipt PaymentReceipt
	err = json.Unmarshal(result.Raw, &paymentReceipt)
	return &paymentReceipt, err

}

// GetSavedOrderInfo Returns saved order info, if any
func (client *Client) GetSavedOrderInfo() (*OrderInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getSavedOrderInfo",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var orderInfo OrderInfo
	err = json.Unmarshal(result.Raw, &orderInfo)
	return &orderInfo, err

}

// DeleteSavedOrderInfo Deletes saved order info
func (client *Client) DeleteSavedOrderInfo() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "deleteSavedOrderInfo",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// DeleteSavedCredentials Deletes saved credentials for all payment provider bots
func (client *Client) DeleteSavedCredentials() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "deleteSavedCredentials",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetSupportUser Returns a user that can be contacted to get support
func (client *Client) GetSupportUser() (*User, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getSupportUser",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var user User
	err = json.Unmarshal(result.Raw, &user)
	return &user, err

}

// GetWallpapers Returns background wallpapers
func (client *Client) GetWallpapers() (*Wallpapers, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getWallpapers",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var wallpapers Wallpapers
	err = json.Unmarshal(result.Raw, &wallpapers)
	return &wallpapers, err

}

// GetLocalizationTargetInfo Returns information about the current localization target. This is an offline request if only_local is true
// @param onlyLocal If true, returns only locally available information without sending network requests
func (client *Client) GetLocalizationTargetInfo(onlyLocal bool) (*LocalizationTargetInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getLocalizationTargetInfo",
		"only_local": onlyLocal,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var localizationTargetInfo LocalizationTargetInfo
	err = json.Unmarshal(result.Raw, &localizationTargetInfo)
	return &localizationTargetInfo, err

}

// GetLanguagePackStrings Returns strings from a language pack in the current localization target by their keys
// @param languagePackId Language pack identifier of the strings to be returned
// @param keys Language pack keys of the strings to be returned; leave empty to request all available strings
func (client *Client) GetLanguagePackStrings(languagePackId string, keys []string) (*LanguagePackStrings, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":            "getLanguagePackStrings",
		"language_pack_id": languagePackId,
		"keys":             keys,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var languagePackStrings LanguagePackStrings
	err = json.Unmarshal(result.Raw, &languagePackStrings)
	return &languagePackStrings, err

}

// SetCustomLanguagePack Adds or changes a custom language pack to the current localization target
// @param info Information about the language pack. Language pack ID must start with 'X', consist only of English letters, digits and hyphens, and must not exceed 64 characters
// @param strings Strings of the new language pack
func (client *Client) SetCustomLanguagePack(info *LanguagePackInfo, strings []LanguagePackString) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "setCustomLanguagePack",
		"info":    info,
		"strings": strings,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// EditCustomLanguagePackInfo Edits information about a custom language pack in the current localization target
// @param info New information about the custom language pack
func (client *Client) EditCustomLanguagePackInfo(info *LanguagePackInfo) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "editCustomLanguagePackInfo",
		"info":  info,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetCustomLanguagePackString Adds, edits or deletes a string in a custom language pack
// @param languagePackId Identifier of a previously added custom language pack in the current localization target
// @param newString New language pack string
func (client *Client) SetCustomLanguagePackString(languagePackId string, newString *LanguagePackString) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":            "setCustomLanguagePackString",
		"language_pack_id": languagePackId,
		"new_string":       newString,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// DeleteLanguagePack Deletes all information about a language pack in the current localization target. The language pack that is currently in use can't be deleted
// @param languagePackId Identifier of the language pack to delete
func (client *Client) DeleteLanguagePack(languagePackId string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":            "deleteLanguagePack",
		"language_pack_id": languagePackId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// RegisterDevice Registers the currently used device for receiving push notifications
// @param deviceToken Device token
// @param otherUserIds List of at most 100 user identifiers of other users currently using the client
func (client *Client) RegisterDevice(deviceToken DeviceToken, otherUserIds []int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":          "registerDevice",
		"device_token":   deviceToken,
		"other_user_ids": otherUserIds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var okDummy Ok
	err = json.Unmarshal(result.Raw, &okDummy)
	return &okDummy, err

}

// GetRecentlyVisitedTMeUrls Returns t.me URLs recently visited by a newly registered user
// @param referrer Google Play referrer to identify the user
func (client *Client) GetRecentlyVisitedTMeUrls(referrer string) (*TMeUrls, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "getRecentlyVisitedTMeUrls",
		"referrer": referrer,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var tMeUrls TMeUrls
	err = json.Unmarshal(result.Raw, &tMeUrls)
	return &tMeUrls, err

}

// SetUserPrivacySettingRules Changes user privacy settings
// @param setting The privacy setting
// @param rules The new privacy rules
func (client *Client) SetUserPrivacySettingRules(setting UserPrivacySetting, rules *UserPrivacySettingRules) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "setUserPrivacySettingRules",
		"setting": setting,
		"rules":   rules,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetUserPrivacySettingRules Returns the current privacy settings
// @param setting The privacy setting
func (client *Client) GetUserPrivacySettingRules(setting UserPrivacySetting) (*UserPrivacySettingRules, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getUserPrivacySettingRules",
		"setting": setting,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var userPrivacySettingRules UserPrivacySettingRules
	err = json.Unmarshal(result.Raw, &userPrivacySettingRules)
	return &userPrivacySettingRules, err

}

// GetOption Returns the value of an option by its name. (Check the list of available options on https://core.telegram.org/tdlib/options.) Can be called before authorization
// @param name The name of the option
func (client *Client) GetOption(name string) (OptionValue, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getOption",
		"name":  name,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	switch OptionValueEnum(result.Data["@type"].(string)) {

	case OptionValueBooleanType:
		var optionValue OptionValueBoolean
		err = json.Unmarshal(result.Raw, &optionValue)
		return &optionValue, err

	case OptionValueEmptyType:
		var optionValue OptionValueEmpty
		err = json.Unmarshal(result.Raw, &optionValue)
		return &optionValue, err

	case OptionValueIntegerType:
		var optionValue OptionValueInteger
		err = json.Unmarshal(result.Raw, &optionValue)
		return &optionValue, err

	case OptionValueStringType:
		var optionValue OptionValueString
		err = json.Unmarshal(result.Raw, &optionValue)
		return &optionValue, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// SetOption Sets the value of an option. (Check the list of available options on https://core.telegram.org/tdlib/options.) Only writable options can be set. Can be called before authorization
// @param name The name of the option
// @param value The new value of the option
func (client *Client) SetOption(name string, value OptionValue) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "setOption",
		"name":  name,
		"value": value,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetAccountTtl Changes the period of inactivity after which the account of the current user will automatically be deleted
// @param ttl New account TTL
func (client *Client) SetAccountTtl(ttl *AccountTtl) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "setAccountTtl",
		"ttl":   ttl,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetAccountTtl Returns the period of inactivity after which the account of the current user will automatically be deleted
func (client *Client) GetAccountTtl() (*AccountTtl, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getAccountTtl",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var accountTtl AccountTtl
	err = json.Unmarshal(result.Raw, &accountTtl)
	return &accountTtl, err

}

// DeleteAccount Deletes the account of the current user, deleting all information associated with the user from the server. The phone number of the account can be used to create a new account. Can be called before authorization when the current authorization state is authorizationStateWaitPassword
// @param reason The reason why the account was deleted; optional
func (client *Client) DeleteAccount(reason string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":  "deleteAccount",
		"reason": reason,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetChatReportSpamState Returns information on whether the current chat can be reported as spam
// @param chatId Chat identifier
func (client *Client) GetChatReportSpamState(chatId int64) (*ChatReportSpamState, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "getChatReportSpamState",
		"chat_id": chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var chatReportSpamState ChatReportSpamState
	err = json.Unmarshal(result.Raw, &chatReportSpamState)
	return &chatReportSpamState, err

}

// ChangeChatReportSpamState Used to let the server know whether a chat is spam or not. Can be used only if ChatReportSpamState.can_report_spam is true. After this request, ChatReportSpamState.can_report_spam becomes false forever
// @param chatId Chat identifier
// @param isSpamChat If true, the chat will be reported as spam; otherwise it will be marked as not spam
func (client *Client) ChangeChatReportSpamState(chatId int64, isSpamChat bool) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "changeChatReportSpamState",
		"chat_id":      chatId,
		"is_spam_chat": isSpamChat,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ReportChat Reports a chat to the Telegram moderators. Supported only for supergroups, channels, or private chats with bots, since other chats can't be checked by moderators
// @param chatId Chat identifier
// @param reason The reason for reporting the chat
// @param messageIds Identifiers of reported messages, if any
func (client *Client) ReportChat(chatId int64, reason ChatReportReason, messageIds []int64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "reportChat",
		"chat_id":     chatId,
		"reason":      reason,
		"message_ids": messageIds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetStorageStatistics Returns storage usage statistics
// @param chatLimit Maximum number of chats with the largest storage usage for which separate statistics should be returned. All other chats will be grouped in entries with chat_id == 0. If the chat info database is not used, the chat_limit is ignored and is always set to 0
func (client *Client) GetStorageStatistics(chatLimit int32) (*StorageStatistics, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "getStorageStatistics",
		"chat_limit": chatLimit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var storageStatistics StorageStatistics
	err = json.Unmarshal(result.Raw, &storageStatistics)
	return &storageStatistics, err

}

// GetStorageStatisticsFast Quickly returns approximate storage usage statistics
func (client *Client) GetStorageStatisticsFast() (*StorageStatisticsFast, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getStorageStatisticsFast",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var storageStatisticsFast StorageStatisticsFast
	err = json.Unmarshal(result.Raw, &storageStatisticsFast)
	return &storageStatisticsFast, err

}

// OptimizeStorage Optimizes storage usage, i.e. deletes some files and returns new storage usage statistics. Secret thumbnails can't be deleted
// @param size Limit on the total size of files after deletion. Pass -1 to use the default limit
// @param ttl Limit on the time that has passed since the last time a file was accessed (or creation time for some filesystems). Pass -1 to use the default limit
// @param count Limit on the total count of files after deletion. Pass -1 to use the default limit
// @param immunityDelay The amount of time after the creation of a file during which it can't be deleted, in seconds. Pass -1 to use the default value
// @param fileTypes If not empty, only files with the given type(s) are considered. By default, all types except thumbnails, profile photos, stickers and wallpapers are deleted
// @param chatIds If not empty, only files from the given chats are considered. Use 0 as chat identifier to delete files not belonging to any chat (e.g., profile photos)
// @param excludeChatIds If not empty, files from the given chats are excluded. Use 0 as chat identifier to exclude all files not belonging to any chat (e.g., profile photos)
// @param chatLimit Same as in getStorageStatistics. Affects only returned statistics
func (client *Client) OptimizeStorage(size int64, ttl int32, count int32, immunityDelay int32, fileTypes []FileType, chatIds []int64, excludeChatIds []int64, chatLimit int32) (*StorageStatistics, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":            "optimizeStorage",
		"size":             size,
		"ttl":              ttl,
		"count":            count,
		"immunity_delay":   immunityDelay,
		"file_types":       fileTypes,
		"chat_ids":         chatIds,
		"exclude_chat_ids": excludeChatIds,
		"chat_limit":       chatLimit,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var storageStatistics StorageStatistics
	err = json.Unmarshal(result.Raw, &storageStatistics)
	return &storageStatistics, err

}

// SetNetworkType Sets the current network type. Can be called before authorization. Calling this method forces all network connections to reopen, mitigating the delay in switching between different networks, so it should be called whenever the network is changed, even if the network type remains the same.
// @param typeParam
func (client *Client) SetNetworkType(typeParam NetworkType) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "setNetworkType",
		"type":  typeParam,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetNetworkStatistics Returns network data usage statistics. Can be called before authorization
// @param onlyCurrent If true, returns only data for the current library launch
func (client *Client) GetNetworkStatistics(onlyCurrent bool) (*NetworkStatistics, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "getNetworkStatistics",
		"only_current": onlyCurrent,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var networkStatistics NetworkStatistics
	err = json.Unmarshal(result.Raw, &networkStatistics)
	return &networkStatistics, err

}

// AddNetworkStatistics Adds the specified data to data usage statistics. Can be called before authorization
// @param entry The network statistics entry with the data to be added to statistics
func (client *Client) AddNetworkStatistics(entry NetworkStatisticsEntry) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "addNetworkStatistics",
		"entry": entry,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// ResetNetworkStatistics Resets all network data usage statistics to zero. Can be called before authorization
func (client *Client) ResetNetworkStatistics() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "resetNetworkStatistics",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetPassportElement Returns one of the available Telegram Passport elements
// @param typeParam Telegram Passport element type
// @param password Password of the current user
func (client *Client) GetPassportElement(typeParam PassportElementType, password string) (PassportElement, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "getPassportElement",
		"type":     typeParam,
		"password": password,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	switch PassportElementEnum(result.Data["@type"].(string)) {

	case PassportElementPersonalDetailsType:
		var passportElement PassportElementPersonalDetails
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementPassportType:
		var passportElement PassportElementPassport
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementDriverLicenseType:
		var passportElement PassportElementDriverLicense
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementIdentityCardType:
		var passportElement PassportElementIdentityCard
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementInternalPassportType:
		var passportElement PassportElementInternalPassport
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementAddressType:
		var passportElement PassportElementAddress
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementUtilityBillType:
		var passportElement PassportElementUtilityBill
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementBankStatementType:
		var passportElement PassportElementBankStatement
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementRentalAgreementType:
		var passportElement PassportElementRentalAgreement
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementPassportRegistrationType:
		var passportElement PassportElementPassportRegistration
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementTemporaryRegistrationType:
		var passportElement PassportElementTemporaryRegistration
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementPhoneNumberType:
		var passportElement PassportElementPhoneNumber
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementEmailAddressType:
		var passportElement PassportElementEmailAddress
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// GetAllPassportElements Returns all available Telegram Passport elements
// @param password Password of the current user
func (client *Client) GetAllPassportElements(password string) (*PassportElements, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "getAllPassportElements",
		"password": password,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var passportElements PassportElements
	err = json.Unmarshal(result.Raw, &passportElements)
	return &passportElements, err

}

// SetPassportElement Adds an element to the user's Telegram Passport. May return an error with a message "PHONE_VERIFICATION_NEEDED" or "EMAIL_VERIFICATION_NEEDED" if the chosen phone number or the chosen email address must be verified first
// @param element Input Telegram Passport element
// @param password Password of the current user
func (client *Client) SetPassportElement(element InputPassportElement, password string) (PassportElement, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "setPassportElement",
		"element":  element,
		"password": password,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	switch PassportElementEnum(result.Data["@type"].(string)) {

	case PassportElementPersonalDetailsType:
		var passportElement PassportElementPersonalDetails
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementPassportType:
		var passportElement PassportElementPassport
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementDriverLicenseType:
		var passportElement PassportElementDriverLicense
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementIdentityCardType:
		var passportElement PassportElementIdentityCard
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementInternalPassportType:
		var passportElement PassportElementInternalPassport
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementAddressType:
		var passportElement PassportElementAddress
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementUtilityBillType:
		var passportElement PassportElementUtilityBill
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementBankStatementType:
		var passportElement PassportElementBankStatement
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementRentalAgreementType:
		var passportElement PassportElementRentalAgreement
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementPassportRegistrationType:
		var passportElement PassportElementPassportRegistration
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementTemporaryRegistrationType:
		var passportElement PassportElementTemporaryRegistration
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementPhoneNumberType:
		var passportElement PassportElementPhoneNumber
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	case PassportElementEmailAddressType:
		var passportElement PassportElementEmailAddress
		err = json.Unmarshal(result.Raw, &passportElement)
		return &passportElement, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}

// DeletePassportElement Deletes a Telegram Passport element
// @param typeParam Element type
func (client *Client) DeletePassportElement(typeParam PassportElementType) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "deletePassportElement",
		"type":  typeParam,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetPassportElementErrors Informs the user that some of the elements in their Telegram Passport contain errors; for bots only. The user will not be able to resend the elements, until the errors are fixed
// @param userId User identifier
// @param errors The errors
func (client *Client) SetPassportElementErrors(userId int32, errors []InputPassportElementError) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "setPassportElementErrors",
		"user_id": userId,
		"errors":  errors,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetPreferredCountryLanguage Returns an IETF language tag of the language preferred in the country, which should be used to fill native fields in Telegram Passport personal details. Returns a 404 error if unknown
// @param countryCode A two-letter ISO 3166-1 alpha-2 country code
func (client *Client) GetPreferredCountryLanguage(countryCode string) (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":        "getPreferredCountryLanguage",
		"country_code": countryCode,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// SendPhoneNumberVerificationCode Sends a code to verify a phone number to be added to a user's Telegram Passport
// @param phoneNumber The phone number of the user, in international format
// @param allowFlashCall Pass true if the authentication code may be sent via flash call to the specified phone number
// @param isCurrentPhoneNumber Pass true if the phone number is used on the current device. Ignored if allow_flash_call is false
func (client *Client) SendPhoneNumberVerificationCode(phoneNumber string, allowFlashCall bool, isCurrentPhoneNumber bool) (*AuthenticationCodeInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                   "sendPhoneNumberVerificationCode",
		"phone_number":            phoneNumber,
		"allow_flash_call":        allowFlashCall,
		"is_current_phone_number": isCurrentPhoneNumber,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var authenticationCodeInfo AuthenticationCodeInfo
	err = json.Unmarshal(result.Raw, &authenticationCodeInfo)
	return &authenticationCodeInfo, err

}

// ResendPhoneNumberVerificationCode Re-sends the code to verify a phone number to be added to a user's Telegram Passport
func (client *Client) ResendPhoneNumberVerificationCode() (*AuthenticationCodeInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "resendPhoneNumberVerificationCode",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var authenticationCodeInfo AuthenticationCodeInfo
	err = json.Unmarshal(result.Raw, &authenticationCodeInfo)
	return &authenticationCodeInfo, err

}

// CheckPhoneNumberVerificationCode Checks the phone number verification code for Telegram Passport
// @param code Verification code
func (client *Client) CheckPhoneNumberVerificationCode(code string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "checkPhoneNumberVerificationCode",
		"code":  code,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SendEmailAddressVerificationCode Sends a code to verify an email address to be added to a user's Telegram Passport
// @param emailAddress Email address
func (client *Client) SendEmailAddressVerificationCode(emailAddress string) (*EmailAddressAuthenticationCodeInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":         "sendEmailAddressVerificationCode",
		"email_address": emailAddress,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var emailAddressAuthenticationCodeInfo EmailAddressAuthenticationCodeInfo
	err = json.Unmarshal(result.Raw, &emailAddressAuthenticationCodeInfo)
	return &emailAddressAuthenticationCodeInfo, err

}

// ResendEmailAddressVerificationCode Re-sends the code to verify an email address to be added to a user's Telegram Passport
func (client *Client) ResendEmailAddressVerificationCode() (*EmailAddressAuthenticationCodeInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "resendEmailAddressVerificationCode",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var emailAddressAuthenticationCodeInfo EmailAddressAuthenticationCodeInfo
	err = json.Unmarshal(result.Raw, &emailAddressAuthenticationCodeInfo)
	return &emailAddressAuthenticationCodeInfo, err

}

// CheckEmailAddressVerificationCode Checks the email address verification code for Telegram Passport
// @param code Verification code
func (client *Client) CheckEmailAddressVerificationCode(code string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "checkEmailAddressVerificationCode",
		"code":  code,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetPassportAuthorizationForm Returns a Telegram Passport authorization form for sharing data with a service
// @param botUserId User identifier of the service's bot
// @param scope Telegram Passport element types requested by the service
// @param publicKey Service's public_key
// @param nonce Authorization form nonce provided by the service
// @param password Password of the current user
func (client *Client) GetPassportAuthorizationForm(botUserId int32, scope string, publicKey string, nonce string, password string) (*PassportAuthorizationForm, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "getPassportAuthorizationForm",
		"bot_user_id": botUserId,
		"scope":       scope,
		"public_key":  publicKey,
		"nonce":       nonce,
		"password":    password,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var passportAuthorizationForm PassportAuthorizationForm
	err = json.Unmarshal(result.Raw, &passportAuthorizationForm)
	return &passportAuthorizationForm, err

}

// SendPassportAuthorizationForm Sends a Telegram Passport authorization form, effectively sharing data with the service
// @param autorizationFormId Authorization form identifier
// @param typeParams Types of Telegram Passport elements chosen by user to complete the authorization form
func (client *Client) SendPassportAuthorizationForm(autorizationFormId int32, typeParams []PassportElementType) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                "sendPassportAuthorizationForm",
		"autorization_form_id": autorizationFormId,
		"types":                typeParams,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SendPhoneNumberConfirmationCode Sends phone number confirmation code. Should be called when user presses "https://t.me/confirmphone?phone=*******&hash=**********" or "tg://confirmphone?phone=*******&hash=**********" link
// @param hash Value of the "hash" parameter from the link
// @param phoneNumber Value of the "phone" parameter from the link
// @param allowFlashCall Pass true if the authentication code may be sent via flash call to the specified phone number
// @param isCurrentPhoneNumber Pass true if the phone number is used on the current device. Ignored if allow_flash_call is false
func (client *Client) SendPhoneNumberConfirmationCode(hash string, phoneNumber string, allowFlashCall bool, isCurrentPhoneNumber bool) (*AuthenticationCodeInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                   "sendPhoneNumberConfirmationCode",
		"hash":                    hash,
		"phone_number":            phoneNumber,
		"allow_flash_call":        allowFlashCall,
		"is_current_phone_number": isCurrentPhoneNumber,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var authenticationCodeInfo AuthenticationCodeInfo
	err = json.Unmarshal(result.Raw, &authenticationCodeInfo)
	return &authenticationCodeInfo, err

}

// ResendPhoneNumberConfirmationCode Resends phone number confirmation code
func (client *Client) ResendPhoneNumberConfirmationCode() (*AuthenticationCodeInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "resendPhoneNumberConfirmationCode",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var authenticationCodeInfo AuthenticationCodeInfo
	err = json.Unmarshal(result.Raw, &authenticationCodeInfo)
	return &authenticationCodeInfo, err

}

// CheckPhoneNumberConfirmationCode Checks phone number confirmation code
// @param code The phone number confirmation code
func (client *Client) CheckPhoneNumberConfirmationCode(code string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "checkPhoneNumberConfirmationCode",
		"code":  code,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetBotUpdatesStatus Informs the server about the number of pending bot updates if they haven't been processed for a long time; for bots only
// @param pendingUpdateCount The number of pending updates
// @param errorMessage The last error message
func (client *Client) SetBotUpdatesStatus(pendingUpdateCount int32, errorMessage string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":                "setBotUpdatesStatus",
		"pending_update_count": pendingUpdateCount,
		"error_message":        errorMessage,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// UploadStickerFile Uploads a PNG image with a sticker; for bots only; returns the uploaded file
// @param userId Sticker file owner
// @param pngSticker PNG image with the sticker; must be up to 512 kB in size and fit in 512x512 square
func (client *Client) UploadStickerFile(userId int32, pngSticker InputFile) (*File, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":       "uploadStickerFile",
		"user_id":     userId,
		"png_sticker": pngSticker,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var file File
	err = json.Unmarshal(result.Raw, &file)
	return &file, err

}

// CreateNewStickerSet Creates a new sticker set; for bots only. Returns the newly created sticker set
// @param userId Sticker set owner
// @param title Sticker set title; 1-64 characters
// @param name Sticker set name. Can contain only English letters, digits and underscores. Must end with *"_by_<bot username>"* (*<bot_username>* is case insensitive); 1-64 characters
// @param isMasks True, if stickers are masks
// @param stickers List of stickers to be added to the set
func (client *Client) CreateNewStickerSet(userId int32, title string, name string, isMasks bool, stickers []InputSticker) (*StickerSet, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "createNewStickerSet",
		"user_id":  userId,
		"title":    title,
		"name":     name,
		"is_masks": isMasks,
		"stickers": stickers,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSet StickerSet
	err = json.Unmarshal(result.Raw, &stickerSet)
	return &stickerSet, err

}

// AddStickerToSet Adds a new sticker to a set; for bots only. Returns the sticker set
// @param userId Sticker set owner
// @param name Sticker set name
// @param sticker Sticker to add to the set
func (client *Client) AddStickerToSet(userId int32, name string, sticker *InputSticker) (*StickerSet, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "addStickerToSet",
		"user_id": userId,
		"name":    name,
		"sticker": sticker,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var stickerSet StickerSet
	err = json.Unmarshal(result.Raw, &stickerSet)
	return &stickerSet, err

}

// SetStickerPositionInSet Changes the position of a sticker in the set to which it belongs; for bots only. The sticker set must have been created by the bot
// @param sticker Sticker
// @param position New position of the sticker in the set, zero-based
func (client *Client) SetStickerPositionInSet(sticker InputFile, position int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "setStickerPositionInSet",
		"sticker":  sticker,
		"position": position,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// RemoveStickerFromSet Removes a sticker from the set to which it belongs; for bots only. The sticker set must have been created by the bot
// @param sticker Sticker
func (client *Client) RemoveStickerFromSet(sticker InputFile) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "removeStickerFromSet",
		"sticker": sticker,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetMapThumbnailFile Returns information about a file with a map thumbnail in PNG format. Only map thumbnail files with size less than 1MB can be downloaded
// @param location Location of the map center
// @param zoom Map zoom level; 13-20
// @param width Map width in pixels before applying scale; 16-1024
// @param height Map height in pixels before applying scale; 16-1024
// @param scale Map scale; 1-3
// @param chatId Identifier of a chat, in which the thumbnail will be shown. Use 0 if unknown
func (client *Client) GetMapThumbnailFile(location *Location, zoom int32, width int32, height int32, scale int32, chatId int64) (*File, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "getMapThumbnailFile",
		"location": location,
		"zoom":     zoom,
		"width":    width,
		"height":   height,
		"scale":    scale,
		"chat_id":  chatId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var file File
	err = json.Unmarshal(result.Raw, &file)
	return &file, err

}

// AcceptTermsOfService Accepts Telegram terms of services
// @param termsOfServiceId Terms of service identifier
func (client *Client) AcceptTermsOfService(termsOfServiceId string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":               "acceptTermsOfService",
		"terms_of_service_id": termsOfServiceId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SendCustomRequest Sends a custom request; for bots only
// @param method The method name
// @param parameters JSON-serialized method parameters
func (client *Client) SendCustomRequest(method string, parameters string) (*CustomRequestResult, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":      "sendCustomRequest",
		"method":     method,
		"parameters": parameters,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var customRequestResult CustomRequestResult
	err = json.Unmarshal(result.Raw, &customRequestResult)
	return &customRequestResult, err

}

// AnswerCustomQuery Answers a custom query; for bots only
// @param customQueryId Identifier of a custom query
// @param data JSON-serialized answer to the query
func (client *Client) AnswerCustomQuery(customQueryId JSONInt64, data string) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":           "answerCustomQuery",
		"custom_query_id": customQueryId,
		"data":            data,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// SetAlarm Succeeds after a specified amount of time has passed. Can be called before authorization. Can be called before initialization
// @param seconds Number of seconds before the function returns
func (client *Client) SetAlarm(seconds float64) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":   "setAlarm",
		"seconds": seconds,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetCountryCode Uses current user IP to found his country. Returns two-letter ISO 3166-1 alpha-2 country code. Can be called before authorization
func (client *Client) GetCountryCode() (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getCountryCode",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// GetInviteText Returns the default text for invitation messages to be used as a placeholder when the current user invites friends to Telegram
func (client *Client) GetInviteText() (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getInviteText",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// GetDeepLinkInfo Returns information about a tg:// deep link. Use "tg://need_update_for_some_feature" or "tg:some_unsupported_feature" for testing. Returns a 404 error for unknown links. Can be called before authorization
// @param link The link
func (client *Client) GetDeepLinkInfo(link string) (*DeepLinkInfo, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getDeepLinkInfo",
		"link":  link,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var deepLinkInfo DeepLinkInfo
	err = json.Unmarshal(result.Raw, &deepLinkInfo)
	return &deepLinkInfo, err

}

// AddProxy Adds a proxy server for network requests. Can be called before authorization
// @param server Proxy server IP address
// @param port Proxy server port
// @param enable True, if the proxy should be enabled
// @param typeParam Proxy type
func (client *Client) AddProxy(server string, port int32, enable bool, typeParam ProxyType) (*Proxy, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":  "addProxy",
		"server": server,
		"port":   port,
		"enable": enable,
		"type":   typeParam,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var proxy Proxy
	err = json.Unmarshal(result.Raw, &proxy)
	return &proxy, err

}

// EditProxy Edits an existing proxy server for network requests. Can be called before authorization
// @param proxyId Proxy identifier
// @param server Proxy server IP address
// @param port Proxy server port
// @param enable True, if the proxy should be enabled
// @param typeParam Proxy type
func (client *Client) EditProxy(proxyId int32, server string, port int32, enable bool, typeParam ProxyType) (*Proxy, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "editProxy",
		"proxy_id": proxyId,
		"server":   server,
		"port":     port,
		"enable":   enable,
		"type":     typeParam,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var proxyDummy Proxy
	err = json.Unmarshal(result.Raw, &proxyDummy)
	return &proxyDummy, err

}

// EnableProxy Enables a proxy. Only one proxy can be enabled at a time. Can be called before authorization
// @param proxyId Proxy identifier
func (client *Client) EnableProxy(proxyId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "enableProxy",
		"proxy_id": proxyId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// DisableProxy Disables the currently enabled proxy. Can be called before authorization
func (client *Client) DisableProxy() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "disableProxy",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// RemoveProxy Removes a proxy server. Can be called before authorization
// @param proxyId Proxy identifier
func (client *Client) RemoveProxy(proxyId int32) (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "removeProxy",
		"proxy_id": proxyId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// GetProxies Returns list of proxies that are currently set up. Can be called before authorization
func (client *Client) GetProxies() (*Proxies, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "getProxies",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var proxies Proxies
	err = json.Unmarshal(result.Raw, &proxies)
	return &proxies, err

}

// GetProxyLink Returns an HTTPS link, which can be used to add a proxy. Available only for SOCKS5 and MTProto proxies. Can be called before authorization
// @param proxyId Proxy identifier
func (client *Client) GetProxyLink(proxyId int32) (*Text, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "getProxyLink",
		"proxy_id": proxyId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var text Text
	err = json.Unmarshal(result.Raw, &text)
	return &text, err

}

// PingProxy Computes time needed to receive a response from a Telegram server through a proxy. Can be called before authorization
// @param proxyId Proxy identifier. Use 0 to ping a Telegram server without a proxy
func (client *Client) PingProxy(proxyId int32) (*Seconds, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type":    "pingProxy",
		"proxy_id": proxyId,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var seconds Seconds
	err = json.Unmarshal(result.Raw, &seconds)
	return &seconds, err

}

// TestCallEmpty Does nothing; for testing only
func (client *Client) TestCallEmpty() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testCallEmpty",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// TestCallString Returns the received string; for testing only
// @param x String to return
func (client *Client) TestCallString(x string) (*TestString, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testCallString",
		"x":     x,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var testString TestString
	err = json.Unmarshal(result.Raw, &testString)
	return &testString, err

}

// TestCallBytes Returns the received bytes; for testing only
// @param x Bytes to return
func (client *Client) TestCallBytes(x []byte) (*TestBytes, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testCallBytes",
		"x":     x,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var testBytes TestBytes
	err = json.Unmarshal(result.Raw, &testBytes)
	return &testBytes, err

}

// TestCallVectorInt Returns the received vector of numbers; for testing only
// @param x Vector of numbers to return
func (client *Client) TestCallVectorInt(x []int32) (*TestVectorInt, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testCallVectorInt",
		"x":     x,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var testVectorInt TestVectorInt
	err = json.Unmarshal(result.Raw, &testVectorInt)
	return &testVectorInt, err

}

// TestCallVectorIntObject Returns the received vector of objects containing a number; for testing only
// @param x Vector of objects to return
func (client *Client) TestCallVectorIntObject(x []TestInt) (*TestVectorIntObject, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testCallVectorIntObject",
		"x":     x,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var testVectorIntObject TestVectorIntObject
	err = json.Unmarshal(result.Raw, &testVectorIntObject)
	return &testVectorIntObject, err

}

// TestCallVectorString For testing only request. Returns the received vector of strings; for testing only
// @param x Vector of strings to return
func (client *Client) TestCallVectorString(x []string) (*TestVectorString, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testCallVectorString",
		"x":     x,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var testVectorString TestVectorString
	err = json.Unmarshal(result.Raw, &testVectorString)
	return &testVectorString, err

}

// TestCallVectorStringObject Returns the received vector of objects containing a string; for testing only
// @param x Vector of objects to return
func (client *Client) TestCallVectorStringObject(x []TestString) (*TestVectorStringObject, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testCallVectorStringObject",
		"x":     x,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var testVectorStringObject TestVectorStringObject
	err = json.Unmarshal(result.Raw, &testVectorStringObject)
	return &testVectorStringObject, err

}

// TestSquareInt Returns the squared received number; for testing only
// @param x Number to square
func (client *Client) TestSquareInt(x int32) (*TestInt, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testSquareInt",
		"x":     x,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var testInt TestInt
	err = json.Unmarshal(result.Raw, &testInt)
	return &testInt, err

}

// TestNetwork Sends a simple network request to the Telegram servers; for testing only
func (client *Client) TestNetwork() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testNetwork",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// TestGetDifference Forces an updates.getDifference call to the Telegram servers; for testing only
func (client *Client) TestGetDifference() (*Ok, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testGetDifference",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	var ok Ok
	err = json.Unmarshal(result.Raw, &ok)
	return &ok, err

}

// TestUseUpdate Does nothing and ensures that the Update object is used; for testing only
func (client *Client) TestUseUpdate() (Update, error) {
	result, err := client.SendAndCatch(UpdateData{
		"@type": "testUseUpdate",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])
	}

	switch UpdateEnum(result.Data["@type"].(string)) {

	case UpdateAuthorizationStateType:
		var update UpdateAuthorizationState
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewMessageType:
		var update UpdateNewMessage
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateMessageSendAcknowledgedType:
		var update UpdateMessageSendAcknowledged
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateMessageSendSucceededType:
		var update UpdateMessageSendSucceeded
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateMessageSendFailedType:
		var update UpdateMessageSendFailed
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateMessageContentType:
		var update UpdateMessageContent
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateMessageEditedType:
		var update UpdateMessageEdited
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateMessageViewsType:
		var update UpdateMessageViews
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateMessageContentOpenedType:
		var update UpdateMessageContentOpened
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateMessageMentionReadType:
		var update UpdateMessageMentionRead
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewChatType:
		var update UpdateNewChat
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatTitleType:
		var update UpdateChatTitle
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatPhotoType:
		var update UpdateChatPhoto
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatLastMessageType:
		var update UpdateChatLastMessage
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatOrderType:
		var update UpdateChatOrder
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatIsPinnedType:
		var update UpdateChatIsPinned
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatIsMarkedAsUnreadType:
		var update UpdateChatIsMarkedAsUnread
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatIsSponsoredType:
		var update UpdateChatIsSponsored
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatDefaultDisableNotificationType:
		var update UpdateChatDefaultDisableNotification
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatReadInboxType:
		var update UpdateChatReadInbox
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatReadOutboxType:
		var update UpdateChatReadOutbox
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatUnreadMentionCountType:
		var update UpdateChatUnreadMentionCount
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatNotificationSettingsType:
		var update UpdateChatNotificationSettings
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateScopeNotificationSettingsType:
		var update UpdateScopeNotificationSettings
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatReplyMarkupType:
		var update UpdateChatReplyMarkup
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateChatDraftMessageType:
		var update UpdateChatDraftMessage
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateDeleteMessagesType:
		var update UpdateDeleteMessages
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateUserChatActionType:
		var update UpdateUserChatAction
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateUserStatusType:
		var update UpdateUserStatus
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateUserType:
		var update UpdateUser
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateBasicGroupType:
		var update UpdateBasicGroup
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateSupergroupType:
		var update UpdateSupergroup
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateSecretChatType:
		var update UpdateSecretChat
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateUserFullInfoType:
		var update UpdateUserFullInfo
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateBasicGroupFullInfoType:
		var update UpdateBasicGroupFullInfo
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateSupergroupFullInfoType:
		var update UpdateSupergroupFullInfo
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateServiceNotificationType:
		var update UpdateServiceNotification
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateFileType:
		var update UpdateFile
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateFileGenerationStartType:
		var update UpdateFileGenerationStart
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateFileGenerationStopType:
		var update UpdateFileGenerationStop
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateCallType:
		var update UpdateCall
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateUserPrivacySettingRulesType:
		var update UpdateUserPrivacySettingRules
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateUnreadMessageCountType:
		var update UpdateUnreadMessageCount
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateUnreadChatCountType:
		var update UpdateUnreadChatCount
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateOptionType:
		var update UpdateOption
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateInstalledStickerSetsType:
		var update UpdateInstalledStickerSets
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateTrendingStickerSetsType:
		var update UpdateTrendingStickerSets
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateRecentStickersType:
		var update UpdateRecentStickers
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateFavoriteStickersType:
		var update UpdateFavoriteStickers
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateSavedAnimationsType:
		var update UpdateSavedAnimations
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateLanguagePackStringsType:
		var update UpdateLanguagePackStrings
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateConnectionStateType:
		var update UpdateConnectionState
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateTermsOfServiceType:
		var update UpdateTermsOfService
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewInlineQueryType:
		var update UpdateNewInlineQuery
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewChosenInlineResultType:
		var update UpdateNewChosenInlineResult
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewCallbackQueryType:
		var update UpdateNewCallbackQuery
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewInlineCallbackQueryType:
		var update UpdateNewInlineCallbackQuery
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewShippingQueryType:
		var update UpdateNewShippingQuery
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewPreCheckoutQueryType:
		var update UpdateNewPreCheckoutQuery
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewCustomEventType:
		var update UpdateNewCustomEvent
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	case UpdateNewCustomQueryType:
		var update UpdateNewCustomQuery
		err = json.Unmarshal(result.Raw, &update)
		return &update, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}
