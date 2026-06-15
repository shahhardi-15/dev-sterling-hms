<template>
  <div class="min-h-screen flex" style="background-color: #f5f7fa">
    <!-- Sidebar -->
    <aside
      class="w-64 bg-white flex flex-col fixed h-full"
      style="border-right: 1px solid #eef0f3"
    >
      <!-- Logo -->
      <div
        class="px-6 py-5 flex items-center gap-3"
        style="border-bottom: 1px solid #eef0f3"
      >
        <div
          class="w-8 h-8 rounded-lg flex items-center justify-center"
          style="background-color: #1e3a5f"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="w-4 h-4 text-white"
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
        <div>
          <p class="text-sm font-bold text-gray-900 tracking-tight">
            Sterling Admin
          </p>
          <p class="text-xs text-gray-400">Hospital Management</p>
        </div>
      </div>

      <!-- Nav links -->
      <nav class="flex-1 px-4 py-5 space-y-0.5 overflow-y-auto">
        <RouterLink
          v-for="link in navLinks"
          :key="link.path"
          :to="link.path"
          class="flex items-center gap-3 px-3 py-2.5 rounded-xl text-sm font-medium transition-all duration-150 relative"
          :class="
            isActive(link.path)
              ? 'bg-blue-50'
              : 'text-gray-500 hover:bg-gray-50'
          "
          :style="isActive(link.path) ? 'color: #1e3a5f;' : ''"
        >
          <span
            v-if="isActive(link.path)"
            class="absolute left-0 top-1/2 -translate-y-1/2 w-0.5 h-5 rounded-r-full"
            style="background-color: #1e3a5f"
          ></span>
          <span
            class="w-5 h-5 flex items-center justify-center flex-shrink-0"
            v-html="link.icon"
          ></span>
          {{ link.label }}
        </RouterLink>
      </nav>

      <!-- User info + logout -->
      <div class="px-4 py-4" style="border-top: 1px solid #eef0f3">
        <div class="flex items-center gap-3 mb-3 px-2">
          <div
            class="w-8 h-8 rounded-full flex items-center justify-center text-white text-xs font-bold flex-shrink-0"
            style="background-color: #1e3a5f"
          >
            {{ userInitial }}
          </div>
          <div class="min-w-0">
            <p class="text-sm font-semibold text-gray-900 truncate">
              {{ auth.userName }}
            </p>
            <p class="text-xs text-gray-400 capitalize">{{ auth.userRole }}</p>
          </div>
        </div>
        <button
          @click="handleLogout"
          class="w-full flex items-center gap-2 px-3 py-2 rounded-xl text-sm text-red-500 hover:bg-red-50 transition font-medium"
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
              d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"
            />
          </svg>
          Sign out
        </button>
      </div>
    </aside>

    <!-- Main content -->
    <main class="ml-64 flex-1 p-8 min-h-screen">
      <RouterView />
    </main>
  </div>
</template>

<script setup>
import { computed } from "vue";
import { RouterLink, RouterView, useRouter, useRoute } from "vue-router";
import { useAuthStore } from "../stores/auth";

const auth = useAuthStore();
const router = useRouter();
const route = useRoute();

const userInitial = computed(
  () => auth.userName?.charAt(0).toUpperCase() || "U",
);

const icons = {
  dashboard: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"/></svg>`,
  users: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z"/></svg>`,
  doctor: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/></svg>`,
  appointments: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"/></svg>`,
  prescription: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"/></svg>`,
  medicine: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M19.428 15.428a2 2 0 00-1.022-.547l-2.387-.477a6 6 0 00-3.86.517l-.318.158a6 6 0 01-3.86.517L6.05 15.21a2 2 0 00-1.806.547M8 4h8l-1 1v5.172a2 2 0 00.586 1.414l5 5c1.26 1.26.367 3.414-1.415 3.414H4.828c-1.782 0-2.674-2.154-1.414-3.414l5-5A2 2 0 009 10.172V5L8 4z"/></svg>`,
  book: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M12 4v16m8-8H4"/></svg>`,
  patient: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z"/></svg>`,
  receptionist: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z"/></svg>`,
  pharmacist: `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"/></svg>`,
};

const allLinks = {
  admin: [
    { label: "Dashboard", path: "/admin", icon: icons.dashboard },
    { label: "Manage Users", path: "/admin/users", icon: icons.users },
    {
      label: "Manage Doctors",
      path: "/admin/users?role=doctor",
      icon: icons.doctor,
    },
    {
      label: "Manage Patients",
      path: "/admin/users?role=patient",
      icon: icons.patient,
    },
    {
      label: "Manage Receptionists",
      path: "/admin/users?role=receptionist",
      icon: icons.receptionist,
    },
    {
      label: "Manage Pharmacists",
      path: "/admin/users?role=pharmacist",
      icon: icons.pharmacist,
    },
    { label: "Billing", path: "/admin/billing", icon: icons.prescription },
  ],
  doctor: [
    { label: "Dashboard", path: "/doctor", icon: icons.dashboard },
    {
      label: "My Appointments",
      path: "/doctor/appointments",
      icon: icons.appointments,
    },
    {
      label: "Write Prescription",
      path: "/doctor/prescriptions/new",
      icon: icons.prescription,
    },
    { label: "Medicines", path: "/doctor/medicines", icon: icons.medicine },
  ],
  patient: [
    { label: "Dashboard", path: "/patient", icon: icons.dashboard },
    { label: "Book Appointment", path: "/patient/book", icon: icons.book },
    {
      label: "My Prescriptions",
      path: "/patient/prescriptions",
      icon: icons.prescription,
    },
  ],
  receptionist: [
    { label: "Dashboard", path: "/receptionist", icon: icons.dashboard },
    {
      label: "Appointments",
      path: "/receptionist/appointments",
      icon: icons.appointments,
    },
  ],
  pharmacist: [
    { label: "Dashboard", path: "/pharmacist", icon: icons.dashboard },
    { label: "Medicines", path: "/pharmacist/medicines", icon: icons.medicine },
  ],
};

const navLinks = computed(() => allLinks[auth.userRole] || []);

function isActive(path) {
  if (path.includes("?")) {
    return route.fullPath === path || route.fullPath.startsWith(path);
  }
  return route.path === path;
}

function handleLogout() {
  auth.logout();
  router.push("/login");
}
</script>
