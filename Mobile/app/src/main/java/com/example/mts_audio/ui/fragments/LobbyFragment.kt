package com.example.mts_audio.ui.fragments

import android.content.ClipData
import android.content.ClipboardManager
import android.content.Context
import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.core.content.ContextCompat.getSystemService
import androidx.fragment.app.viewModels
import androidx.lifecycle.Observer
import androidx.navigation.fragment.navArgs
import androidx.recyclerview.widget.LinearLayoutManager
import com.example.mts_audio.MessageRecyclerViewAdapter
import com.example.mts_audio.R
import com.example.mts_audio.databinding.FragmentLobbyBinding
import com.example.mts_audio.ui.model.MessageItem
import com.example.mts_audio.ui.viewmodels.LobbyViewModel
import dagger.hilt.android.AndroidEntryPoint
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.GlobalScope
import kotlinx.coroutines.launch

@AndroidEntryPoint
class LobbyFragment : Fragment() {

    companion object {
        fun newInstance() = LobbyFragment()
    }

    private lateinit var binding: FragmentLobbyBinding
    private val viewModel: LobbyViewModel by viewModels()

    private var lobbyId: String = ""
    private var isOwner: Boolean = false


    private val chatData: MutableList<MessageItem> = mutableListOf()


    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        binding = FragmentLobbyBinding.inflate(layoutInflater)
        return binding.root
    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        lobbyId = arguments?.getString("id").toString()
        isOwner = arguments?.getBoolean("isOwner")!!


        binding.linkCopy.text = lobbyId
        binding.copyButton.setOnClickListener {
            val clipboardManager = requireContext().getSystemService(Context.CLIPBOARD_SERVICE) as ClipboardManager

            val clipData = ClipData.newPlainText("Label", lobbyId)

            clipboardManager.setPrimaryClip(clipData)
        }

        binding.recyclerViewMessage.layoutManager = LinearLayoutManager(requireContext())
        setRoom(lobbyId)

        binding.send.setOnClickListener {
            sendMessage(binding.message.text.toString())
        }

        viewModel.lobbyMessages.observe(viewLifecycleOwner, Observer {
            val lobbyMessage = it ?: return@Observer

            chatData.add(lobbyMessage)
            binding.recyclerViewMessage.adapter = MessageRecyclerViewAdapter(chatData)
        })
    }

    private fun sendMessage(message: String) {
        GlobalScope.launch(Dispatchers.Main) {
            viewModel.sendMessageToChat(message)
        }
    }

    override fun onDestroy() {
        super.onDestroy()
        viewModel.closeConnection()
    }

    private fun setRoom(roomId: String) {
        GlobalScope.launch(Dispatchers.Main) {
            viewModel.setRoom(roomId)
        }
    }

}