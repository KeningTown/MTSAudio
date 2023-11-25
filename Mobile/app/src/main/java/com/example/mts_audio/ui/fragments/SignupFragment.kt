package com.example.mts_audio.ui.fragments

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import android.util.Log
import android.util.Patterns
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.view.inputmethod.EditorInfo
import android.widget.Toast
import androidx.annotation.StringRes
import androidx.fragment.app.viewModels
import androidx.lifecycle.LiveData
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.Observer
import androidx.navigation.fragment.findNavController
import com.example.mts_audio.R
import com.example.mts_audio.data.local.User
import com.example.mts_audio.data.model.Result
import com.example.mts_audio.data.remote.auth.AuthResult
import com.example.mts_audio.databinding.FragmentLoginBinding
import com.example.mts_audio.databinding.FragmentSignupBinding
import com.example.mts_audio.ui.model.AuthFormState
import com.example.mts_audio.ui.viewmodels.SignupViewModel
import com.example.mts_audio.ui.viewmodels.UserViewModel
import dagger.hilt.android.AndroidEntryPoint
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.GlobalScope
import kotlinx.coroutines.launch

@AndroidEntryPoint
class SignupFragment : Fragment() {

    companion object {
        fun newInstance() = SignupFragment()
    }

    private val userViewModel: UserViewModel by viewModels()
    private val viewModel: SignupViewModel by viewModels()

    private lateinit var binding: FragmentSignupBinding

    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        binding = FragmentSignupBinding.inflate(inflater)


        userViewModel.user.observe(viewLifecycleOwner, Observer {
            if (it != null){
                findNavController().navigate(R.id.action_signupFragment_to_mainFragment)
            }
        })

        return binding.root
    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)




        binding.loginButton.setOnClickListener {
            findNavController().navigateUp()
        }

        viewModel.signupState.observe(viewLifecycleOwner, Observer {
            val loginState = it ?: return@Observer


            if (loginState.usernameError != null) {
            } else {
            }

            if (loginState.passwordError != null) {
            } else {
            }
        })

        viewModel.signupResult.observe(viewLifecycleOwner, Observer {
            val loginResult = it ?: return@Observer

            if (loginResult.error != null) {
                showLoginFailed(loginResult.error)
            }
            if (loginResult.success != null) {
                updateUiWithUser(loginResult.success.user.username)
                userViewModel.saveUser(
                    User(
                        loginResult.success.user.id,
                        loginResult.toString(),
                        loginResult.success.access_token,
                    )
                )

            }
        })

        binding.userPasswordInput.apply {
            afterTextChanged {
                viewModel.signupDataChanged(
                    binding.userNameInput.text.toString(),
                    binding.userPasswordInput.text.toString()
                )
            }

            setOnEditorActionListener { _, actionId, _ ->
                when (actionId) {
                    EditorInfo.IME_ACTION_DONE ->
                        signup(
                            binding.userNameInput.text.toString(),
                            binding.userPasswordInput.text.toString(),
                        )
                }
                false
            }

            binding.signupButton.setOnClickListener {
                signup(
                    binding.userNameInput.text.toString(),
                    binding.userPasswordInput.text.toString(),
                )
            }
        }
    }


    private fun signup(username: String, password: String) {
        GlobalScope.launch(Dispatchers.Main) {
            viewModel.signup(password, username)
        }
    }

    private fun updateUiWithUser(name: String) {
        val welcome = getString(R.string.welcome)
        // TODO : initiate successful logged in experience
        Toast.makeText(
            context,
            "$welcome $name",
            Toast.LENGTH_LONG
        ).show()
    }

    private fun showLoginFailed(@StringRes errorString: Int) {
        Toast.makeText(context, errorString, Toast.LENGTH_SHORT).show()
    }
}