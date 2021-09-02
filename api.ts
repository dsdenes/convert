export interface SMSConversationMessage {
  direction: string
  isSuccessful: boolean
  isAutomatic: boolean
  timestamp: string
  text: string
}

export interface SMSConversation {
  phoneNumber: string
  lastInteractionAt: string
  lastMessageText: string
  lastMessageFailed: boolean
}

export type GetConversationsResponse = {
  data: SMSConversation[]
}

export type GetConversationMessagesResponse = {
  data: SMSConversationMessage[]
}

export interface SendMessageToConversationRequest {
  text: string
}

export interface SendMessageToConversationResponse {
  error?: any
}
