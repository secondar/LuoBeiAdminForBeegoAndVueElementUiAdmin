import CheckPermissions from '@/utils/permission'
import store from '@/store'

function checkPermission(el, binding) {
  const { value } = binding
  try {
    if (!CheckPermissions(store.getters.menus, value)) {
      el.parentNode && el.parentNode.removeChild(el)
    }
  } catch (err) {
    //  el.parentNode && el.parentNode.removeChild(el)
  }
}

export default {
  inserted(el, binding) {
    checkPermission(el, binding)
  },
  update(el, binding) {
    checkPermission(el, binding)
  }
}
