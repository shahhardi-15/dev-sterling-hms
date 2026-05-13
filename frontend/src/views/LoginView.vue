<template>
  <div class="min-h-screen bg-gray-50 flex flex-col">
    <nav class="px-8 py-5">
      <span class="font-bold text-lg tracking-tight" style="color: #1e3a5f">
        Sterling Hospital Management System
      </span>
    </nav>

    <div class="flex-1 flex items-center justify-center px-6 py-10">
      <div
        class="w-full max-w-5xl bg-white rounded-2xl shadow-md border border-gray-100 flex overflow-hidden min-h-[520px]"
      >
        <div
          class="w-1/2 bg-slate-100 p-10 flex flex-col justify-between border-r border-gray-200"
        >
          <div>
            <div
              class="w-12 h-12 rounded-xl flex items-center justify-center mb-8"
              style="background-color: #1e3a5f"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="w-6 h-6 text-white"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M9 12h6m-3-3v6m-7 4h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"
                />
              </svg>
            </div>
            <h1 class="text-3xl font-bold text-gray-900 leading-tight mb-4">
              Clinical Excellence<br />at Your Fingertips.
            </h1>
            <p class="text-gray-500 text-sm leading-relaxed">
              Access the Sterling ecosystem with secure, high-performance
              administrative tools designed for modern healthcare professionals.
            </p>
          </div>
          <div>
            <div class="border-t border-gray-200 mb-4"></div>
            <div class="flex items-center gap-3">
              <div class="w-2 h-2 bg-green-500 rounded-full"></div>
              <span class="text-xs text-gray-500 tracking-widest uppercase"
                >System Status: Optimal</span
              >
            </div>
          </div>
        </div>

        <div class="w-1/2 p-10 flex flex-col justify-center">
          <h2 class="text-2xl font-bold text-gray-900 mb-1">Welcome Back</h2>
          <p class="text-gray-500 text-sm mb-8">
            Please enter your credentials to access your dashboard.
          </p>

          <div
            v-if="error"
            class="bg-red-50 border border-red-200 text-red-600 text-sm px-4 py-3 rounded-lg mb-6"
          >
            {{ error }}
          </div>

          <div class="mb-5">
            <label
              class="block text-xs font-semibold text-gray-600 tracking-widest uppercase mb-2"
              >Email ID</label
            >
            <div class="relative">
              <span
                class="absolute left-4 top-1/2 -translate-y-1/2 text-gray-400"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="w-4 h-4"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"
                  />
                </svg>
              </span>
              <input
                v-model="email"
                type="email"
                placeholder="doctor.name@sterling.com"
                class="w-full pl-11 pr-4 py-3 bg-gray-50 border border-gray-200 rounded-xl text-sm text-gray-800 placeholder-gray-400 focus:outline-none focus:border-blue-500 focus:bg-white transition"
              />
            </div>
          </div>

          <div class="mb-5">
            <div class="flex items-center justify-between mb-2">
              <label
                class="block text-xs font-semibold text-gray-600 tracking-widest uppercase"
                >Password</label
              >
              <a href="#" class="text-xs hover:underline" style="color: #1e3a5f"
                >Contact Admin to reset password</a
              >
            </div>
            <div class="relative">
              <span
                class="absolute left-4 top-1/2 -translate-y-1/2 text-gray-400"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="w-4 h-4"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"
                  />
                </svg>
              </span>
              <input
                :type="showPassword ? 'text' : 'password'"
                v-model="password"
                placeholder="••••••••••"
                class="w-full pl-11 pr-10 py-3 bg-gray-50 border border-gray-200 rounded-xl text-sm text-gray-800 placeholder-gray-400 focus:outline-none focus:border-blue-500 focus:bg-white transition"
              />
              <button
                type="button"
                @click="showPassword = !showPassword"
                class="absolute right-4 top-1/2 -translate-y-1/2 text-gray-400 hover:text-gray-600"
              >
                <svg
                  v-if="showPassword"
                  xmlns="http://www.w3.org/2000/svg"
                  class="w-4 h-4"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.88 9.88l-3.29-3.29m7.532 7.532l3.29 3.29M3 3l3.59 3.59m0 0A9.953 9.953 0 0112 5c4.478 0 8.268 2.943 9.543 7a10.025 10.025 0 01-4.132 4.411m0 0L21 21"
                  />
                </svg>
                <svg
                  v-else
                  xmlns="http://www.w3.org/2000/svg"
                  class="w-4 h-4"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
                  />
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"
                  />
                </svg>
              </button>
            </div>
          </div>

          <div class="flex items-center gap-2 mb-8">
            <input
              type="checkbox"
              id="remember"
              v-model="remember"
              class="w-4 h-4 rounded border-gray-300 text-blue-600"
            />
            <label for="remember" class="text-sm text-gray-600"
              >Remember me</label
            >
          </div>

          <button
            @click="handleLogin"
            :disabled="loading"
            class="w-full text-white font-semibold py-3 rounded-xl flex items-center justify-center gap-2 transition disabled:opacity-60"
            style="background-color: #1e3a5f"
          >
            <span>{{ loading ? "Signing in..." : "Sign In to Sterling" }}</span>
            <svg
              v-if="!loading"
              xmlns="http://www.w3.org/2000/svg"
              class="w-4 h-4"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M17 8l4 4m0 0l-4 4m4-4H3"
              />
            </svg>
          </button>
        </div>
      </div>
    </div>

    <footer
      class="px-8 py-5 flex items-center justify-between text-xs text-gray-400"
    >
      <span>© 2026 Sterling Clinical Excellence. All rights reserved.</span>
      <div class="flex gap-6">
        <a href="#" class="hover:text-gray-600">Privacy Policy</a>
        <a href="#" class="hover:text-gray-600">Terms of Service</a>
        <a href="#" class="hover:text-gray-600">Help Center</a>
        <a href="#" class="hover:text-gray-600">Contact</a>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import { useAuthStore } from "../stores/auth";

const router = useRouter();
const auth = useAuthStore();

const email = ref("");
const password = ref("");
const remember = ref(false);
const loading = ref(false);
const error = ref("");
const showPassword = ref(false);

async function handleLogin() {
  if (!email.value || !password.value) {
    error.value = "Please enter your email and password.";
    return;
  }
  loading.value = true;
  error.value = "";
  try {
    const user = await auth.login(email.value, password.value);
    router.push(`/${user.role}`);
  } catch (err) {
    error.value = err.response?.data?.error || "Invalid email or password.";
  } finally {
    loading.value = false;
  }
}
</script>
