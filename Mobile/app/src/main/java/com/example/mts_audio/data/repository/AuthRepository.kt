package com.example.mts_audio.data.repository

import android.util.Log
import com.example.mts_audio.data.remote.auth.AuthApi
import com.example.mts_audio.data.remote.auth.AuthResponse
import com.example.mts_audio.data.remote.auth.UserLoginRequest
import com.example.mts_audio.data.model.Result
import com.example.mts_audio.data.remote.auth.UserRegistrationRequest
import kotlinx.coroutines.CoroutineDispatcher
import kotlinx.coroutines.withContext
import java.io.IOException

class AuthRepository(
    private val authDataSource: AuthApi,
    private val ioDispatcher: CoroutineDispatcher,
) {

    suspend fun login(email: String, password: String): Result<AuthResponse> {
        try {
            return Result.Success(
                withContext(ioDispatcher) {
                    val response = authDataSource.loginUser(UserLoginRequest(email, password))
                    response.await()
                }
            )
        } catch (e: Exception) {
            Log.d("TAG", e.message.toString())
            return Result.Error(IOException("Error logging in", e))
        }
    }

    suspend fun signup(email: String, password: String, username: String): Result<AuthResponse> {
        try {
            return Result.Success(
                withContext(ioDispatcher) {
                    val response = authDataSource.registerUser(UserRegistrationRequest(email, password, username))
                    response.await()
                }
            )
        } catch (e: Exception) {
            Log.d("TAG", e.message.toString())
            return Result.Error(IOException("Error signup in", e))
        }

    }

    suspend fun refresh(token: String): Result<AuthResponse>{
        try {
            return Result.Success(
                withContext(ioDispatcher) {
                    val response = authDataSource.refresh(token)
                    response.await()
                }
            )
        } catch (e: Exception) {
            Log.d("TAG", e.message.toString())
            return Result.Error(IOException("Error refresh", e))
        }
    }

}