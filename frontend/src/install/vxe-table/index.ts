// import { VxeUI, VxeIcon, VxeLoading } from 'vxe-pc-ui'
import { VxeTable, VxeColumn } from 'vxe-table';

// 导入主题变量，也可以重写主题变量
import 'vxe-table/styles/cssvar.scss'
// import 'vxe-pc-ui/styles/cssvar.scss'



export const vxeTableInstall = (app: any) => {
    app.use(VxeTable)
    app.use(VxeColumn)
}