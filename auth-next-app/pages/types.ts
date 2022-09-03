export interface ApiResponse<T> {
    statusCode: number
    message: string
    payload: T
}

export interface GetFlowIdPayload {
    eula: string
    flowId: string
}

export interface PostAuthenticatePayload {
    token: string
}

export interface PostAuthenticateRequest {
    flowId: string
    signature: string
}
