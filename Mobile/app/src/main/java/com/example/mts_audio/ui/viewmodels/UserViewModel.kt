package com.example.mts_audio.ui.viewmodels

import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.example.mts_audio.data.local.User
import com.example.mts_audio.data.repository.LocalUserRepository
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.launch
import javax.inject.Inject

@HiltViewModel
class UserViewModel @Inject constructor(
    private val userRepository: LocalUserRepository,
): ViewModel(){

    private val _user = MutableLiveData<User?>()
    val user: LiveData<User?> get() = _user

    init {
        viewModelScope.launch {
            userRepository.userFlow.collect { user ->
                _user.value = user
            }
        }
    }

    fun saveUser(user: User) {
        userRepository.saveUser(user)
    }

    fun deleteUser() {
        userRepository.clearUserData()
    }

}