package com.example.mts_audio.ui.fragments

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import android.text.Editable
import android.text.TextWatcher
import android.util.Log
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.view.inputmethod.EditorInfo
import android.widget.EditText
import android.widget.Toast
import androidx.annotation.StringRes
import androidx.fragment.app.viewModels
import androidx.lifecycle.Observer
import androidx.navigation.fragment.findNavController
import com.example.mts_audio.R
import com.example.mts_audio.data.local.User
import com.example.mts_audio.databinding.FragmentLoginBinding
import com.example.mts_audio.ui.viewmodels.LoginViewModel
import com.example.mts_audio.ui.viewmodels.UserViewModel
import dagger.hilt.android.AndroidEntryPoint
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.GlobalScope
import kotlinx.coroutines.launch


@AndroidEntryPoint
class LoginFragment : Fragment() {

    companion object {
        fun newInstance() = LoginFragment()
    }

    private val userViewModel: UserViewModel by viewModels()
    private val viewModel: LoginViewModel by viewModels()

    private lateinit var binding: FragmentLoginBinding

    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        binding = FragmentLoginBinding.inflate(layoutInflater)

        userViewModel.user.observe(viewLifecycleOwner) { user ->
            if (user != null) run {
                findNavController().navigate(R.id.action_loginFragment_to_mainFragment)
            }
        }

        return binding.root
    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)


        binding.signupButton.setOnClickListener {
            findNavController().navigate(R.id.action_loginFragment_to_signupFragment)
        }

        viewModel.loginState.observe(viewLifecycleOwner, Observer {
            val loginState = it ?: return@Observer


            if (loginState.usernameError != null) {
            } else {
            }

            if (loginState.passwordError != null) {
            } else {
            }
        })

        viewModel.loginResult.observe(viewLifecycleOwner, Observer {
            val loginResult = it ?: return@Observer

            if (loginResult.error != null) {
                showLoginFailed(loginResult.error)
            }
            if (loginResult.success != null) {
                updateUiWithUser(loginResult.success.user.username)
                Log.d("TAG", "${loginResult.success}")
                userViewModel.saveUser(
                    User(
                        loginResult.success.user.id,
                        loginResult.success.user.username,
                        binding.userPasswordInput.toString(),
                        loginResult.success.access_token,
                    )
                )
            }
        })

        binding.userNameInput.afterTextChanged {
            viewModel.loginDataChanged(
                binding.userNameInput.text.toString(),
                binding.userPasswordInput.text.toString()
            )
        }

        binding.userPasswordInput.apply {
            afterTextChanged {
                viewModel.loginDataChanged(
                    binding.userNameInput.text.toString(),
                    binding.userPasswordInput.text.toString()
                )
            }

            setOnEditorActionListener { _, actionId, _ ->
                when (actionId) {
                    EditorInfo.IME_ACTION_DONE ->
                        login(
                            binding.userNameInput.text.toString(),
                            binding.userPasswordInput.text.toString(),
                        )
                }
                false
            }

            binding.loginButton.setOnClickListener {
                login(
                    binding.userNameInput.text.toString(),
                    binding.userPasswordInput.text.toString(),
                    )
            }
        }
    }

    private fun login(username: String, password: String) {
        GlobalScope.launch(Dispatchers.Main) {
            viewModel.login(username, password)
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

fun EditText.afterTextChanged(afterTextChanged: (String) -> Unit) {
    this.addTextChangedListener(object : TextWatcher {
        override fun afterTextChanged(editable: Editable?) {
            afterTextChanged.invoke(editable.toString())
        }

        override fun beforeTextChanged(s: CharSequence, start: Int, count: Int, after: Int) {}

        override fun onTextChanged(s: CharSequence, start: Int, before: Int, count: Int) {}
    })
}