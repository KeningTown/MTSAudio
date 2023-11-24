package com.example.mts_audio.ui

import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import com.example.mts_audio.R
import dagger.hilt.android.AndroidEntryPoint

@AndroidEntryPoint
class MainActivity : AppCompatActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)
    }
}