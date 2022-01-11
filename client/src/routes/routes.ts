export const AuthedRoutes = {
  index: '/',
}

export const UnAuthedRoutes = {
  index: '/',
}

export const routes = { ...AuthedRoutes, ...UnAuthedRoutes }
