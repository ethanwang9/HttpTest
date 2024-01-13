import {resolve} from "path"
import {type ConfigEnv, loadEnv, type UserConfigExport} from "vite"
import vue from '@vitejs/plugin-vue'
import Icons from 'unplugin-icons/vite'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import IconsResolver from 'unplugin-icons/resolver'
import {ElementPlusResolver} from 'unplugin-vue-components/resolvers'
import ElementPlus from 'unplugin-element-plus/vite'
import viteLegacyPlugin from "@vitejs/plugin-legacy";

export default (configEnv: ConfigEnv): UserConfigExport => {
    const env = loadEnv(configEnv.mode, process.cwd())

    let config: UserConfigExport = {
        base: "./",
        resolve: {
            alias: {
                '@': resolve(__dirname, './src'),
                '~': resolve(__dirname, './'),
            }
        },
        plugins: [
            vue(),
            AutoImport({
                imports: ['vue', 'vue-router'],
                resolvers: [
                    ElementPlusResolver(),
                    IconsResolver({
                        prefix: 'Icon',
                    }),
                ],
            }),
            Components({
                resolvers: [
                    IconsResolver({
                        enabledCollections: ['ep'],
                    }),
                    ElementPlusResolver(),
                ],
            }),
            Icons({
                autoInstall: true,
                compiler: "vue3",
            }),
            ElementPlus({
                useSource: true,
            }),
            viteLegacyPlugin({
                targets: ['Android > 39', 'Chrome >= 60', 'Safari >= 10.1', 'iOS >= 10.3', 'Firefox >= 54', 'Edge >= 15'],
            }),
        ],
        build: {
            chunkSizeWarningLimit: 2000, // 消除打包大小超过500kb警告
            minify: 'terser', // 是否进行压缩,boolean | 'terser' | 'esbuild',默认使用terser
            manifest: false, // 是否产出manifest.json
            sourcemap: false, // 是否产出sourcemap.json
            outDir: 'dist', // 产出目录
            terserOptions: {
                mangle: true, // 开启变量名混淆
                compress: {
                    keep_infinity: true, // 防止 Infinity 被压缩成 1/0，这可能会导致 Chrome 上的性能问题
                    drop_console: true, // 生产环境去除 console
                    drop_debugger: true, // 生产环境去除 debugger
                },
                format: {
                    comments: false, // 删除注释
                },
            },
            rollupOptions: {
                output: {
                    // 用于从入口点创建的块的打包输出格式[name]表示文件名,[hash]表示该文件内容hash值
                    entryFileNames: "js/[name].[hash].js",
                    // 用于命名代码拆分时创建的共享块的输出命名
                    chunkFileNames: "js/[name].[hash].js",
                    // 用于输出静态资源的命名，[ext]表示文件扩展名
                    assetFileNames: (assetInfo: any) => {
                        const info = assetInfo.name.split(".");
                        let extType = info[info.length - 1];
                        if (
                            /\.(mp4|webm|ogg|mp3|wav|flac|aac)(\?.*)?$/i.test(assetInfo.name)
                        ) {
                            extType = "media";
                        } else if (/\.(png|jpe?g|gif|svg)(\?.*)?$/.test(assetInfo.name)) {
                            extType = "img";
                        } else if (/\.(woff2?|eot|ttf|otf)(\?.*)?$/i.test(assetInfo.name)) {
                            extType = "fonts";
                        }
                        return `${extType}/[name].[hash].[ext]`;
                    },
                },
            },
        }
    }

    if (configEnv.mode === "development") {
        config.server = {
            host: true,
            port: parseInt(env.VITE_WEB_PORT),
            proxy: {
                [env.VITE_SERVER_PATH]: {
                    target: `http://${env.VITE_SERVER_ADDR}:${env.VITE_SERVER_PORT}/`,
                    changeOrigin: true,
                    rewrite: (path) => path.replace(new RegExp('^' + env.VITE_SERVER_PATH), ''),
                },
            },
        }
    }

    return config
}