import { Router, RouterModule } from '@angular/router';
import { Component, NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormBuilder, Validators, ReactiveFormsModule } from '@angular/forms';
import { CardModule } from 'primeng/card';
import { ButtonModule } from 'primeng/button';
import { InputTextModule } from 'primeng/inputtext';
import { AuthService } from '../../services/auth.service';
import { User } from '../../interfaces/auth';
import { MessageService } from 'primeng/api';
import { ToastModule } from 'primeng/toast';


@Component({
  selector: 'app-register',
  standalone: true,
  imports: [
    CommonModule,
    CardModule,
    InputTextModule,
    ReactiveFormsModule,
    ButtonModule,
    RouterModule
  ],
  templateUrl: './register.component.html',
  styleUrl: './register.component.css'
})
export class RegisterComponent {

  registerForm = this.fb.group({
    email: ['', [Validators.required, Validators.email]],
    password: ['', Validators.required]
  })

  constructor (
    private fb: FormBuilder,
    private authService: AuthService,
    private messageService: MessageService,
    private router: Router
    ) {  }

  get email() {
    return this.registerForm.controls['email']
  }

  get password() {
    return this.registerForm.controls['password']
  }

  submitDetails() {
    const postData = { ...this.registerForm.value };
    this.authService.registerUser(postData as User).subscribe(
    response => {
      console.log(response);
      this.messageService.add({ severity: 'success', summary: 'Success', detail: 'Register succsessfully' });
      this.router.navigate(['/login'])
    },
    error => {
      this.messageService.add({ severity: 'error', summary: 'Error', detail: 'Something went wrong' });
    }
    )
  }
}
