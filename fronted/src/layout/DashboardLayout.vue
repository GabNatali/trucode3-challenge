<script setup lang="ts">
import { onMounted } from "vue";
import { initFlowbite } from 'flowbite'
import { RouterView } from "vue-router";
import { useAuthStore } from "@/stores/auth.store";


const authStore = useAuthStore();
onMounted(async() => {
  initFlowbite();
});

</script>
  <template>
  <nav class="bg-white border border-b-gray-300">
      <div class="flex flex-wrap items-center justify-between mx-auto px-4 md:px-12 pb-4 md:py-6">
          <div class="flex items-center order-2 space-x-1 md:space-x-2 rtl:space-x-reverse w-full justify-end md:w-auto">
            <button id="mega-menu-dropdown-button" data-dropdown-toggle="mega-menu-dropdown" class="flex items-center justify-center gap-2">
                <span>{{ authStore.userData.email }}</span>
                <img class="w-8 h-8 rounded-full" src="https://flowbite.com/docs/images/people/profile-picture-3.jpg" alt="user photo">
              </button>
              <div id="mega-menu-dropdown" class="absolute z-10 grid w-48 text-sm bg-white border border-gray-100 rounded-lg shadow-md hidden">
                          <div class="p-4 text-gray-900 ">
                              <button
                                @click.prevent="authStore.logout()"
                                class="space-y-4 flex items-center" aria-labelledby="mega-menu-dropdown-button" >
                                    Logout
                              </button>
                          </div>
              </div>
          </div>
          <div id="mega-menu" class="items-center justify-between w-full md:flex md:w-auto sm:order-1">
              <ul class="flex flex-col mt-4 font-medium md:flex-row md:mt-0 md:space-x-8 rtl:space-x-reverse">
                  <li>
                    <RouterLink
                    active-class="text-blue-500"
                    :to="{ name: 'home' }" class="block py-2 px-3 text-gray-900 border-b border-gray-100 hover:bg-gray-50 md:hover:bg-transparent md:border-0 md:hover:text-blue-600 md:p-0">
                    Home
                  </RouterLink>
                  </li>
                  <li>
                    <RouterLink
                    active-class="text-blue-500"
                    :to="{ name: 'user' }" class="block py-2 px-3 text-gray-900 border-b border-gray-100 hover:bg-gray-50 md:hover:bg-transparent md:border-0 md:hover:text-blue-600 md:p-0">Configuraciones</RouterLink>
                  </li>
              </ul>
          </div>
      </div>
  </nav>
  <main class="flex h-full"><RouterView /></main>

</template>

