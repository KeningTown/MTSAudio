package com.example.mts_audio.ui.fragments

import android.content.ClipData
import android.content.ClipboardManager
import android.content.Context
import android.media.MediaPlayer
import android.os.Bundle
import android.util.Log
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.fragment.app.viewModels
import androidx.lifecycle.Observer
import androidx.recyclerview.widget.LinearLayoutManager
import com.example.mts_audio.MessageRecyclerViewAdapter
import com.example.mts_audio.databinding.FragmentLobbyBinding
import com.example.mts_audio.ui.model.MessageItem
import com.example.mts_audio.ui.viewmodels.LobbyViewModel
import dagger.hilt.android.AndroidEntryPoint
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.GlobalScope
import kotlinx.coroutines.launch
import java.io.FileOutputStream
import java.io.IOException

@AndroidEntryPoint
class LobbyFragment : Fragment() {

    companion object {
        fun newInstance() = LobbyFragment()
    }

    private lateinit var binding: FragmentLobbyBinding
    private val viewModel: LobbyViewModel by viewModels()

    private var lobbyId: String = ""
    private var isOwner: Boolean = false


    private val mediaPlayer: MediaPlayer = MediaPlayer()
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

        viewModel.lobbyMusic.observe(viewLifecycleOwner, Observer{
            val lobbyMusic = it ?: return@Observer

            cashByteArrayToMp3(lobbyMusic)
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

    private fun playAudio(music: String) {

        try {
            if (mediaPlayer.isPlaying) {
                mediaPlayer.stop()
            }
            try {
                mediaPlayer.setDataSource(music)
                mediaPlayer.prepare()
                mediaPlayer.start()
            } catch (e: IOException) {
                e.printStackTrace()
            } catch (e: IllegalArgumentException) {
                e.printStackTrace()
            } catch (e: SecurityException) {
                e.printStackTrace()
            } catch (e: IllegalStateException) {
                e.printStackTrace()
            }
        } catch (e: Exception) {
            e.printStackTrace()
        }

    }

    private fun cashByteArrayToMp3(lobbyMusic: ByteArray) {

        if (lobbyMusic.isEmpty()) {
            Log.e("LobbyFragment", "Received empty lobbyMusic ByteArray")
            return
        }
        val tempFile = java.io.File.createTempFile("tempAudio", ".mp3", requireContext().cacheDir)
        Log.d("MediaPlayer", "Temp file path: ${tempFile.absolutePath}")
        Log.d("MediaPlayer", "Byte array size: ${lobbyMusic.size}")


        val outputStream = FileOutputStream(tempFile)
        outputStream.write(lobbyMusic)
        outputStream.close()

        Log.d("MediaPlayer", "Setting data source from file: ${tempFile.absolutePath}")

        playAudio(tempFile.absolutePath)
    }

}