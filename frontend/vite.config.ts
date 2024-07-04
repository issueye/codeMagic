import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import Icons from 'unplugin-icons/vite'
import Components from 'unplugin-vue-components/vite'
import { lazyImport, VxeResolver } from 'vite-plugin-lazy-import'


// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    Icons({}),
    

    Components({}),
    lazyImport({
      resolvers: [
        VxeResolver({
          libraryName: 'vxe-table'
        }),
        // VxeResolver({
        //   libraryName: 'vxe-pc-ui',
        // })
      ]
    })
  ],
  resolve: {
    alias: {
      '@': './src'
    },
  },
})
