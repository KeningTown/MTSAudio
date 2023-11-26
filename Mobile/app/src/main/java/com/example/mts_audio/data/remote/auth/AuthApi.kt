package com.example.mts_audio.data.remote.auth

import kotlinx.coroutines.Deferred
import retrofit2.Response
import retrofit2.http.Body
import retrofit2.http.GET
import retrofit2.http.Header
import retrofit2.http.Headers
import retrofit2.http.POST

interface AuthApi {
    @POST("api/Account/SignUp")
    @Headers("Content-Type: application/json")
    fun registerUser(@Body request: UserRegistrationRequest): Deferred<AuthResponse>

    @POST("api/Account/SignIn")
    @Headers("Content-Type: application/json")
    fun loginUser(@Body request: UserLoginRequest): Deferred<AuthResponse>

    @POST("api/Account/SignOut")
    @Headers("Content-Type: application/json")
    fun logoutUser(): Deferred<Void>

    @GET("api/Account/RefreshTokens")
    @Headers("Content-Type: application/json")
    fun refreshToken(
        @Header("Authorization") token: String,
    ): Response<AuthResponse>

    @GET("api/Account/RefreshTokens")
    @Headers("Content-Type: application/json")
    fun refresh(
        @Header("Authorization") token: String,
    ): Deferred<AuthResponse>
}