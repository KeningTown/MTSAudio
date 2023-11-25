package com.example.mts_audio.data.repository

import android.content.SharedPreferences
import com.example.mts_audio.data.local.User
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow
import javax.inject.Inject

class LocalUserRepository @Inject constructor(
    private val sharedPreferences: SharedPreferences
) {

    private val KEY_ACCESS_TOKEN = "access_token"
    private val KEY_USER_ID = "user_id"

    //TODO delete user password from sharedPreferences
    private val KEY_USER_PASSWORD = "user_password"

    private val editor: SharedPreferences.Editor = sharedPreferences.edit()

    private val _userFlow = MutableStateFlow<User?>(null)
    val userFlow: Flow<User?> get() = _userFlow.asStateFlow()

    init {
        _userFlow.value = getUser()
    }

    fun saveUser(user: User) {
        with(editor) {
            putString(KEY_USER_ID, user.userid)
            putString(KEY_USER_PASSWORD, user.userPassword)
            putString(KEY_ACCESS_TOKEN, user.accessToken)
            apply()
        }

        _userFlow.value = getUser()
    }

    fun getUser(): User? {
        val userId = getUserId()
        val userPassword = getUserPassword()
        val accessToken = getAccessToken()


        return if (userId.isNullOrEmpty() || userPassword.isNullOrEmpty() || accessToken.isNullOrEmpty()) {
            null
        } else User(
            userid = userId,
            userPassword = userPassword,
            accessToken = accessToken,
        )
    }

    fun saveJWToken(accessToken: String) {
        with(editor) {
            putString(KEY_ACCESS_TOKEN, accessToken)
            apply()
        }

        _userFlow.value = getUser()
    }

    fun getUserId(): String? {
        return sharedPreferences.getString(KEY_USER_ID, null)
    }

    fun getUserPassword(): String? {
        return sharedPreferences.getString(KEY_USER_PASSWORD, null)
    }


    fun getAccessToken(): String? {
        return sharedPreferences.getString(KEY_ACCESS_TOKEN, null)
    }

    fun clearUserData() {
        with(editor) {
            remove(KEY_ACCESS_TOKEN)
            remove(KEY_USER_ID)
            remove(KEY_USER_PASSWORD)
            apply()
        }
        _userFlow.value = null
    }
}