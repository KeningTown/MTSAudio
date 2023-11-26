package com.example.mts_audio.ui.fragments

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.Button
import android.widget.EditText
import androidx.fragment.app.viewModels
import androidx.lifecycle.Observer
import androidx.navigation.fragment.findNavController
import com.example.mts_audio.R
import com.example.mts_audio.data.remote.websocket.WebSocketManager
import com.example.mts_audio.ui.viewmodels.MainViewModel
import com.example.mts_audio.ui.viewmodels.UserViewModel
import dagger.hilt.android.AndroidEntryPoint
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.GlobalScope
import kotlinx.coroutines.launch

@AndroidEntryPoint
class MainFragment : Fragment() {

    companion object {
        fun newInstance() = MainFragment()
    }

    private val viewModel: MainViewModel by viewModels()
    private val userModel: UserViewModel by viewModels()


    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        return inflater.inflate(R.layout.fragment_main, container, false)
    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        userModel.user.observe(viewLifecycleOwner, Observer {
            if (it == null){
                findNavController().navigate(R.id.loginFragment)
            }
        })

        viewModel.roomResult.observe(viewLifecycleOwner, Observer {
            val result = it?:return@Observer

            if (result.success !=null){
                val bundle = Bundle()
                bundle.putString("id", result.success.roomId)
                bundle.putBoolean("isOwner", true)
                findNavController().navigate(R.id.action_mainFragment_to_lobbyFragment, bundle)
            }
        })

        var newLobbyButton = view.findViewById<Button>(R.id.create_lobby)
        var signInLobbyButton = view.findViewById<Button>(R.id.enter_lobby)

        var inputId = view.findViewById<EditText>(R.id.lobby_id)

        newLobbyButton.setOnClickListener {
            getRoom()
        }

        signInLobbyButton.setOnClickListener {
            val bundle = Bundle()
            bundle.putString("id", inputId.text.toString())
            bundle.putBoolean("isOwner", true)
            findNavController().navigate(R.id.action_mainFragment_to_lobbyFragment, bundle)
        }
    }

    private fun getRoom() {
        GlobalScope.launch(Dispatchers.Main) {
            viewModel.getRoom()
        }
    }

}