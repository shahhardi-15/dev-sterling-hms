import { createRouter, createWebHistory } from "vue-router";
import { useAuthStore } from "../stores/auth";

const routes = [
  {
    path: "/",
    redirect: "/login",
  },
  {
    path: "/login",
    name: "Login",
    component: () => import("../views/LoginView.vue"),
    meta: { requiresGuest: true },
  },
  {
    path: "/unauthorized",
    name: "Unauthorized",
    component: () => import("../views/UnauthorizedView.vue"),
  },

  // Admin routes
  {
    path: "/admin",
    component: () => import("../layouts/DashboardLayout.vue"),
    meta: { requiresAuth: true, roles: ["admin"] },
    children: [
      {
        path: "",
        name: "AdminDashboard",
        component: () => import("../views/admin/AdminDashboard.vue"),
      },
      {
        path: "users",
        name: "ManageUsers",
        component: () => import("../views/admin/ManageUsers.vue"),
      },
      {
        path: "billing",
        name: "AdminBilling",
        component: () => import("../views/admin/AdminBilling.vue"),
      },
    ],
  },

  // Doctor routes
  {
    path: "/doctor",
    component: () => import("../layouts/DashboardLayout.vue"),
    meta: { requiresAuth: true, roles: ["doctor"] },
    children: [
      {
        path: "",
        name: "DoctorDashboard",
        component: () => import("../views/doctor/DoctorDashboard.vue"),
      },
      {
        path: "appointments",
        name: "DoctorAppointments",
        component: () => import("../views/doctor/MyAppointments.vue"),
      },
      {
        path: "prescriptions/new",
        name: "WritePrescription",
        component: () => import("../views/doctor/WritePrescription.vue"),
      },
      // Under doctor children, after WritePrescription route, add:
      {
        path: "prescriptions",
        name: "DoctorPrescriptions",
        component: () => import("../views/doctor/DoctorPrescriptions.vue"),
      },
      {
        path: "medicines",
        name: "DoctorMedicines",
        component: () => import("../views/doctor/DoctorMedicines.vue"),
      },
    ],
  },

  // Patient routes

  {
    path: "/patient",
    component: () => import("../layouts/DashboardLayout.vue"),
    meta: { requiresAuth: true, roles: ["patient"] },
    children: [
      {
        path: "",
        name: "PatientDashboard",
        component: () => import("../views/patient/PatientDashboard.vue"),
      },
      {
        path: "book",
        name: "BookAppointment",
        component: () => import("../views/patient/BookAppointment.vue"),
      },
      {
        path: "prescriptions",
        name: "MyPrescriptions",
        component: () => import("../views/patient/MyPrescriptions.vue"),
      },
    ],
  },

  // Receptionist routes
  {
    path: "/receptionist",
    component: () => import("../layouts/DashboardLayout.vue"),
    meta: { requiresAuth: true, roles: ["receptionist"] },
    children: [
      {
        path: "",
        name: "ReceptionistDashboard",
        component: () =>
          import("../views/receptionist/ReceptionistDashboard.vue"),
      },
      {
        path: "appointments",
        name: "ReceptionistAppointments",
        component: () =>
          import("../views/receptionist/ReceptionistAppointments.vue"),
      },
    ],
  },

  // Pharmacist routes
  {
    path: "/pharmacist",
    component: () => import("../layouts/DashboardLayout.vue"),
    meta: { requiresAuth: true, roles: ["pharmacist"] },
    children: [
      {
        path: "",
        name: "PharmacistDashboard",
        component: () => import("../views/pharmacist/PharmacistDashboard.vue"),
      },
      // Under pharmacist children, after PharmacistDashboard route, add:
      {
        path: "medicines",
        name: "PharmacistMedicines",
        component: () => import("../views/pharmacist/PharmacistMedicines.vue"),
      },
    ],
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

// Navigation guards
router.beforeEach((to, from, next) => {
  const token = localStorage.getItem("token");
  const user = JSON.parse(localStorage.getItem("user") || "null");
  const isLoggedIn = !!token;
  const userRole = user?.role || null;

  if (to.meta.requiresAuth && !isLoggedIn) {
    return next("/login");
  }

  if (to.meta.requiresGuest && isLoggedIn) {
    return next(`/${userRole}`);
  }

  if (to.meta.roles && !to.meta.roles.includes(userRole)) {
    return next("/unauthorized");
  }

  next();
});
export default router;
