import store from '@/store'

/**
 * @param {Array} value
 * @returns {Boolean}
 * @example see @/views/permission/directive.vue
 */
export default function checkPermission(value) {
  return CheckPermissions(store.getters.menus, value)
}
// 递归检查按钮权限
export function CheckPermissions(menus, routers, isNext = false) {
  if (routers == "") routers += ""
  for (let i in menus) {
    let arrTmp = menus[i]
    if (arrTmp.characteristic.replace(/\s*/g, "") == routers.replace(/\s*/g, "")) {
      isNext = true
      break
    } else {
      if (menus[i].children != null && menus[i].children.length > 0) {
        isNext = CheckPermissions(menus[i].children, routers, isNext);
      }
    }
  }
  return isNext;
}