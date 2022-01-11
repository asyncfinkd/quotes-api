import lazy from 'react-lazy-with-preload'
import { routes } from './routes'

export const authRoutesData = [
  {
    title: 'Index',
    path: routes.index,
    component: lazy(() => import('pages/index')),
  },
]

export const unAuthRoutesData = [
  {
    title: 'Index',
    path: routes.index,
    component: lazy(() => import('pages/index')),
  },
]

export const authRoutes = [...authRoutesData]
export const unAuthRoutes = [...unAuthRoutesData]
