package com.example.mts_audio.ui.viewmodels

import android.util.Log
import android.util.Patterns
import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel
import com.example.mts_audio.R
import com.example.mts_audio.data.model.Result
import com.example.mts_audio.data.remote.auth.AuthResult
import com.example.mts_audio.data.repository.AuthRepository
import com.example.mts_audio.data.repository.LocalUserRepository
import com.example.mts_audio.ui.model.AuthFormState
import dagger.hilt.android.lifecycle.HiltViewModel
import javax.inject.Inject

@HiltViewModel
class LoginViewModel @Inject constructor(
    private val authRepository: AuthRepository,
    private val localUserRepository: LocalUserRepository,
) : ViewModel() {

    private val _loginState  = MutableLiveData<AuthFormState>()
    val loginState: LiveData<AuthFormState> = _loginState

    private val _loginResult = MutableLiveData<AuthResult>()
    val loginResult: LiveData<AuthResult> = _loginResult

    suspend fun login(email: String, password: String) {
        val result = authRepository.login(email, password)

        if (result is Result.Success) {
            _loginResult.value = AuthResult(success = result.data)
            Log.d("TAG", "access token${result.data.access_token}")
        } else {
            _loginResult.value = AuthResult(error = R.string.error_string)
        }
    }

    fun loginDataChanged(username: String, password: String) {
        if (!isUserNameValid(username)) {
            _loginState.value = AuthFormState(usernameError = R.string.invalid_username)
        } else if (!isPasswordValid(password)) {
            _loginState.value = AuthFormState(passwordError = R.string.invalid_password)
        } else {
            _loginState.value = AuthFormState(isDataValid = true)
        }
    }

    private fun isUserNameValid(username: String): Boolean {
        return if (username.contains('@')) {
            Patterns.EMAIL_ADDRESS.matcher(username).matches()
        } else {
            username.isNotBlank()
        }
    }

    private fun isPasswordValid(password: String): Boolean {
        return password.length > 5
    }
}