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
import com.example.mts_audio.ui.model.AuthFormState
import dagger.hilt.android.lifecycle.HiltViewModel
import javax.inject.Inject

@HiltViewModel
class SignupViewModel @Inject constructor(
    private val authRepository: AuthRepository,
) : ViewModel() {

    private val _signupState  = MutableLiveData<AuthFormState>()
    val signupState: LiveData<AuthFormState> = _signupState

    private val _signupResult = MutableLiveData<AuthResult>()
    val signupResult: LiveData<AuthResult> = _signupResult

    suspend fun signup(password: String, username: String) {
        val result = authRepository.signup(password, username)

        if (result is Result.Success) {
            _signupResult.value = AuthResult(success = result.data)
            Log.d("TAG", "access token${result.data.access_token}")
        } else {
            _signupResult.value = AuthResult(error = R.string.error_string)
        }
    }

    fun signupDataChanged(username: String, password: String) {
        if (!isUserNameValid(username)) {
            _signupState.value = AuthFormState(usernameError = R.string.invalid_username)
        } else if (!isPasswordValid(password)) {
            _signupState.value = AuthFormState(passwordError = R.string.invalid_password)
        } else {
            _signupState.value = AuthFormState(isDataValid = true)
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