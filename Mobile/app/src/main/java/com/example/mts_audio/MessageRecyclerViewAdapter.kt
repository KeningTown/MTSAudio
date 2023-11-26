package com.example.mts_audio

import android.graphics.Color
import android.graphics.Typeface
import android.text.SpannableString
import android.text.style.ForegroundColorSpan
import android.text.style.StyleSpan
import androidx.recyclerview.widget.RecyclerView
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView

import com.example.mts_audio.placeholder.PlaceholderContent.PlaceholderItem
import com.example.mts_audio.databinding.FragmentMessageBinding
import com.example.mts_audio.ui.model.MessageItem

/**
 * [RecyclerView.Adapter] that can display a [PlaceholderItem].
 * TODO: Replace the implementation with code for your data type.
 */
class MessageRecyclerViewAdapter(
    private val values: List<MessageItem>
) : RecyclerView.Adapter<MessageRecyclerViewAdapter.ViewHolder>() {

    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): ViewHolder {

        return ViewHolder(
            FragmentMessageBinding.inflate(
                LayoutInflater.from(parent.context),
                parent,
                false
            )
        )

    }

    override fun onBindViewHolder(holder: ViewHolder, position: Int) {
        val item = values[position]

        val username = item.message.username
        val msg = item.message.msg

        val fullText = "${username}: ${msg}"

        val spannableString = SpannableString(fullText)

        val usernameColor = Color.BLUE
        spannableString.setSpan(ForegroundColorSpan(usernameColor), 0, username.length, 0)

        holder.messageItem.text = spannableString
    }

    override fun getItemCount(): Int = values.size

    inner class ViewHolder(binding: FragmentMessageBinding) :
        RecyclerView.ViewHolder(binding.root) {
        val messageItem: TextView = binding.message

        override fun toString(): String {
            return super.toString() + " '" + messageItem.text + "'"
        }
    }

}