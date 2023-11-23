package com.example.mts_audio.common

import android.util.Log
import com.example.mts_audio.data.repository.LocalUserRepository
import com.jakewharton.retrofit2.adapter.kotlin.coroutines.CoroutineCallAdapterFactory
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.runBlocking
import kotlinx.coroutines.withContext
import okhttp3.Authenticator
import okhttp3.OkHttpClient
import okhttp3.Request
import okhttp3.Response
import okhttp3.Route
import okhttp3.logging.HttpLoggingInterceptor
import retrofit2.Retrofit
import retrofit2.converter.gson.GsonConverterFactory
import java.io.IOException
import javax.inject.Inject
import com.example.mts_audio.data.model.Result
import com.example.mts_audio.data.remote.auth.AuthApi
import com.example.mts_audio.data.remote.auth.AuthResponse

class AuthAuthenticator @Inject constructor(
    private val userRepository: LocalUserRepository
) : Authenticator {


    override fun authenticate(route: Route?, response: Response): Request? {
        return runBlocking {
            val token = userRepository.getRefreshToken()
            when (val newToken = getNewToken(token)) {
                is Result.Success -> {
                    newToken.data.apply {
                        userRepository.saveJWToken(refreshToken, accessToken)
                    }
                    response.request.newBuilder()
                        .header("Authorization", "Bearer ${newToken.data.accessToken}")
                        .build()
                }
                is Result.Error -> {
                    userRepository.clearUserData()
                    null
                }
            }
        }
    }

    private suspend fun getNewToken(refreshToken: String?): Result<AuthResponse> {
        val loggingInterceptor = HttpLoggingInterceptor()
        loggingInterceptor.level = HttpLoggingInterceptor.Level.BODY
        val okHttpClient = OkHttpClient.Builder()
            .addInterceptor(loggingInterceptor)
            .build()
        val retrofit = Retrofit.Builder()
            .baseUrl("")
            .addConverterFactory(GsonConverterFactory.create())
            .addCallAdapterFactory(CoroutineCallAdapterFactory())
            .client(okHttpClient)
            .build()
        val service = retrofit.create(AuthApi::class.java)
        return try {
            Result.Success(
                withContext(Dispatchers.IO) {
                    service.refresh(refreshToken!!).await()
                }
            )
        } catch (e: Exception) {
            Log.d("TAG", e.message.toString())
            Result.Error(IOException("Error refresh", e))
        }
    }
}