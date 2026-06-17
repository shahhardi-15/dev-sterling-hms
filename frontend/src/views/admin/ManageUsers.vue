<template>
  <div>
    <!-- Header -->
    <div class="mb-8 flex items-center justify-between">
      <div>
        <h1 class="text-3xl font-bold text-gray-900">{{ pageTitle }}</h1>
        <p class="text-gray-400 text-sm mt-1">Create and manage system users</p>
      </div>
      <button
        @click="showModal = true"
        class="text-white text-sm font-semibold px-5 py-2.5 rounded-xl transition flex items-center gap-2"
        style="background-color: #1e3a5f"
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
            d="M12 4v16m8-8H4"
          />
        </svg>
        Add User
      </button>
    </div>

    <!-- Filter Tabs -->
    <div class="flex gap-2 mb-6 flex-wrap">
      <button
        v-for="f in filters"
        :key="f.value"
        @click="activeFilter = f.value"
        class="px-4 py-2 rounded-xl text-sm font-medium transition"
        :class="
          activeFilter === f.value
            ? 'text-white'
            : 'bg-white border border-gray-200 text-gray-500 hover:bg-gray-50'
        "
        :style="activeFilter === f.value ? 'background-color: #1e3a5f;' : ''"
      >
        {{ f.label }}
        <span class="ml-1 text-xs opacity-70">
          ({{
            f.value === "all"
              ? users.length
              : users.filter((u) => u.role === f.value).length
          }})
        </span>
      </button>
    </div>

    <!-- Users table -->
    <div class="bg-white rounded-2xl border border-gray-100 shadow-sm">
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead>
            <tr class="border-b border-gray-100">
              <th
                class="text-left px-6 py-3 text-xs font-semibold text-gray-400 uppercase tracking-widest"
              >
                Name
              </th>
              <th
                class="text-left px-6 py-3 text-xs font-semibold text-gray-400 uppercase tracking-widest"
              >
                Email
              </th>
              <th
                class="text-left px-6 py-3 text-xs font-semibold text-gray-400 uppercase tracking-widest"
              >
                Role
              </th>
              <th
                class="text-left px-6 py-3 text-xs font-semibold text-gray-400 uppercase tracking-widest"
              >
                Status
              </th>
              <th
                class="text-left px-6 py-3 text-xs font-semibold text-gray-400 uppercase tracking-widest"
              >
                Actions
              </th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="user in filteredUsers"
              :key="user.id"
              class="border-b border-gray-50 hover:bg-gray-50 transition"
            >
              <td class="px-6 py-4">
                <div class="flex items-center gap-3">
                  <div
                    class="w-8 h-8 rounded-full flex items-center justify-center text-white text-xs font-bold flex-shrink-0"
                    style="background-color: #1e3a5f"
                  >
                    {{ user.full_name?.charAt(0) }}
                  </div>
                  <p class="text-sm font-medium text-gray-900">
                    {{ user.full_name }}
                  </p>
                </div>
              </td>
              <td class="px-6 py-4 text-sm text-gray-500">{{ user.email }}</td>
              <td class="px-6 py-4">
                <span
                  class="px-2.5 py-1 rounded-full text-xs font-semibold capitalize"
                  :class="{
                    'bg-blue-50 text-blue-600': user.role === 'doctor',
                    'bg-green-50 text-green-600': user.role === 'patient',
                    'bg-purple-50 text-purple-600': user.role === 'admin',
                    'bg-orange-50 text-orange-600':
                      user.role === 'receptionist',
                    'bg-pink-50 text-pink-600': user.role === 'pharmacist',
                  }"
                >
                  {{ user.role }}
                </span>
              </td>
              <td class="px-6 py-4">
                <span
                  class="px-2.5 py-1 rounded-full text-xs font-semibold"
                  :class="
                    user.is_active
                      ? 'bg-green-50 text-green-600'
                      : 'bg-red-50 text-red-600'
                  "
                >
                  {{ user.is_active ? "Active" : "Inactive" }}
                </span>
              </td>
              <td class="px-6 py-4">
                <div class="flex gap-3">
                  <button
                    @click="openResetModal(user)"
                    class="text-xs font-medium hover:underline"
                    style="color: #1e3a5f"
                  >
                    Reset Password
                  </button>
                  <button
                    @click="deleteUser(user.id)"
                    class="text-xs text-red-500 hover:text-red-600 font-medium hover:underline"
                  >
                    Delete
                  </button>
                </div>
              </td>
            </tr>
            <tr v-if="filteredUsers.length === 0">
              <td
                colspan="5"
                class="px-6 py-8 text-center text-sm text-gray-400"
              >
                No users found
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Add User Modal -->
    <div
      v-if="showModal"
      class="fixed inset-0 bg-black/40 flex items-center justify-center z-50"
    >
      <div class="bg-white rounded-2xl p-8 w-full max-w-md shadow-xl">
        <h2 class="text-xl font-bold text-gray-900 mb-1">Add New User</h2>
        <p class="text-sm text-gray-400 mb-6">
          Fill in the details to create a new user
        </p>
        <div
          v-if="formError"
          class="bg-red-50 border border-red-200 text-red-600 text-sm px-4 py-3 rounded-xl mb-4"
        >
          {{ formError }}
        </div>
        <div class="space-y-4">
          <div>
            <label
              class="block text-xs font-semibold text-gray-500 uppercase tracking-widest mb-1.5"
              >Full Name</label
            >
            <input
              v-model="form.full_name"
              type="text"
              placeholder="e.g. Dr. John Smith"
              class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-xl text-sm focus:outline-none transition"
            />
          </div>
          <div>
            <label
              class="block text-xs font-semibold text-gray-500 uppercase tracking-widest mb-1.5"
              >Email</label
            >
            <input
              v-model="form.email"
              type="email"
              placeholder="john@sterling.com"
              class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-xl text-sm focus:outline-none transition"
            />
          </div>
          <div>
            <label
              class="block text-xs font-semibold text-gray-500 uppercase tracking-widest mb-1.5"
              >Password</label
            >
            <input
              v-model="form.password"
              type="password"
              placeholder="Min. 8 characters"
              class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-xl text-sm focus:outline-none transition"
            />
          </div>
          <div>
            <label
              class="block text-xs font-semibold text-gray-500 uppercase tracking-widest mb-1.5"
              >Role</label
            >
            <select
              v-model="form.role"
              class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-xl text-sm focus:outline-none transition"
            >
              <option value="">Select role</option>
              <option value="admin">Admin</option>
              <option value="doctor">Doctor</option>
              <option value="patient">Patient</option>
              <option value="receptionist">Receptionist</option>
              <option value="pharmacist">Pharmacist</option>
            </select>
          </div>
        </div>
        <div class="flex gap-3 mt-6">
          <button
            @click="showModal = false"
            class="flex-1 px-4 py-3 border border-gray-200 rounded-xl text-sm font-medium text-gray-600 hover:bg-gray-50 transition"
          >
            Cancel
          </button>
          <button
            @click="createUser"
            :disabled="creating"
            class="flex-1 px-4 py-3 text-white rounded-xl text-sm font-semibold transition disabled:opacity-60"
            style="background-color: #1e3a5f"
          >
            {{ creating ? "Creating..." : "Create User" }}
          </button>
        </div>
      </div>
    </div>

    <!-- Reset Password Modal -->
    <div
      v-if="showResetModal"
      class="fixed inset-0 bg-black/40 flex items-center justify-center z-50"
    >
      <div class="bg-white rounded-2xl p-8 w-full max-w-md shadow-xl">
        <h2 class="text-xl font-bold text-gray-900 mb-1">Reset Password</h2>
        <p class="text-sm text-gray-400 mb-6">
          Set a new password for
          <span class="font-semibold text-gray-700">{{
            resetUser?.full_name
          }}</span>
        </p>
        <div
          v-if="resetError"
          class="bg-red-50 border border-red-200 text-red-600 text-sm px-4 py-3 rounded-xl mb-4"
        >
          {{ resetError }}
        </div>
        <div
          v-if="resetSuccess"
          class="bg-green-50 border border-green-200 text-green-600 text-sm px-4 py-3 rounded-xl mb-4"
        >
          Password reset successfully!
        </div>
        <div class="mb-6">
          <label
            class="block text-xs font-semibold text-gray-500 uppercase tracking-widest mb-1.5"
            >New Password</label
          >
          <div class="relative">
            <input
              :type="showNewPassword ? 'text' : 'password'"
              v-model="newPassword"
              placeholder="Enter new password (min 6 characters)"
              class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-xl text-sm focus:outline-none transition pr-10"
            />
            <button
              type="button"
              @click="showNewPassword = !showNewPassword"
              class="absolute right-4 top-1/2 -translate-y-1/2 text-gray-400 hover:text-gray-600"
            >
              <svg
                v-if="showNewPassword"
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
        <div class="flex gap-3">
          <button
            @click="showResetModal = false"
            class="flex-1 px-4 py-3 border border-gray-200 rounded-xl text-sm font-medium text-gray-600 hover:bg-gray-50 transition"
          >
            Cancel
          </button>
          <button
            @click="resetPassword"
            :disabled="resetting"
            class="flex-1 px-4 py-3 text-white rounded-xl text-sm font-semibold transition disabled:opacity-60"
            style="background-color: #1e3a5f"
          >
            {{ resetting ? "Resetting..." : "Reset Password" }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from "vue";
import { useRoute } from "vue-router";
import api from "../../api/axios";

const route = useRoute();
const users = ref([]);
const showModal = ref(false);
const creating = ref(false);
const formError = ref("");
const activeFilter = ref(route.query.role || "all");

const form = ref({ full_name: "", email: "", password: "", role: "" });

const filters = [
  { label: "All", value: "all" },
  { label: "Doctors", value: "doctor" },
  { label: "Patients", value: "patient" },
  { label: "Receptionists", value: "receptionist" },
  { label: "Pharmacists", value: "pharmacist" },
  { label: "Admins", value: "admin" },
];

const filteredUsers = computed(() => {
  if (activeFilter.value === "all") return users.value;
  return users.value.filter((u) => u.role === activeFilter.value);
});

const pageTitle = computed(() => {
  const f = filters.find((f) => f.value === activeFilter.value);
  return f
    ? f.value === "all"
      ? "Manage Users"
      : `Manage ${f.label}`
    : "Manage Users";
});

watch(
  () => route.query.role,
  (newRole) => {
    activeFilter.value = newRole || "all";
  },
);

const showResetModal = ref(false);
const resetUser = ref(null);
const newPassword = ref("");
const resetting = ref(false);
const resetError = ref("");
const resetSuccess = ref(false);
const showNewPassword = ref(false);

onMounted(async () => {
  await loadUsers();
});

async function loadUsers() {
  try {
    const res = await api.get("/users");
    users.value = res.data.data;
  } catch (err) {
    console.error(err);
  }
}

async function createUser() {
  if (
    !form.value.full_name ||
    !form.value.email ||
    !form.value.password ||
    !form.value.role
  ) {
    formError.value = "Please fill in all fields.";
    return;
  }
  creating.value = true;
  formError.value = "";
  try {
    await api.post("/users", form.value);
    showModal.value = false;
    form.value = { full_name: "", email: "", password: "", role: "" };
    await loadUsers();
  } catch (err) {
    const raw = err.response?.data?.error || "";
    if (raw.includes("Password") && raw.includes("min")) {
      formError.value = "Password must be at least 8 characters long.";
    } else if (raw.includes("Email") || raw.includes("email")) {
      formError.value = "Please enter a valid email address.";
    } else if (raw.includes("unique") || raw.includes("duplicate")) {
      formError.value = "A user with this email already exists.";
    } else if (raw) {
      formError.value = raw;
    } else {
      formError.value = "Failed to create user. Please try again.";
    }
  } finally {
    creating.value = false;
  }
}

function openResetModal(user) {
  resetUser.value = user;
  newPassword.value = "";
  resetError.value = "";
  resetSuccess.value = false;
  showNewPassword.value = false;
  showResetModal.value = true;
}

async function resetPassword() {
  if (!newPassword.value || newPassword.value.length < 6) {
    resetError.value = "Password must be at least 6 characters.";
    return;
  }
  resetting.value = true;
  resetError.value = "";
  resetSuccess.value = false;
  try {
    await api.put(`/users/${resetUser.value.id}`, {
      full_name: resetUser.value.full_name,
      is_active: resetUser.value.is_active,
      password: newPassword.value,
    });
    resetSuccess.value = true;
    newPassword.value = "";
    setTimeout(() => {
      showResetModal.value = false;
    }, 1500);
  } catch (err) {
    resetError.value = err.response?.data?.error || "Failed to reset password.";
  } finally {
    resetting.value = false;
  }
}

async function deleteUser(id) {
  if (!confirm("Are you sure you want to delete this user?")) return;
  try {
    await api.delete(`/users/${id}`);
    await loadUsers();
  } catch (err) {
    console.error(err);
  }
}
</script>
