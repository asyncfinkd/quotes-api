import React, { Suspense } from 'react'
import ReactDOM from 'react-dom'
import './index.css'
import reportWebVitals from './reportWebVitals'
import NProgress from 'nprogress'
import 'nprogress/nprogress.css'
import { Toaster } from 'react-hot-toast'
import { BrowserRouter, Navigate, Route, Routes } from 'react-router-dom'
import { NavigationScroll } from './shared/nav-scroll'
import { authRoutes, unAuthRoutesData } from 'routes/route-generator'

const LazyLoad = () => {
  React.useEffect(() => {
    NProgress.configure({ showSpinner: false })
    NProgress.start()

    return () => {
      NProgress.done()
    }
  })

  return <></>
}

const Pages = () => {
  if (!localStorage.getItem('token')) {
    return (
      <Routes>
        {unAuthRoutesData.map((el) => {
          return (
            <Route key={el.path} path={el.path} element={<el.component />} />
          )
        })}
        <Route path="*" element={<Navigate to="/" />} />
      </Routes>
    )
  }

  return (
    <Routes>
      {authRoutes.map((el) => {
        return <Route key={el.path} path={el.path} element={<el.component />} />
      })}
      <Route path="*" element={<Navigate to="/" />} />
    </Routes>
  )
}

ReactDOM.render(
  <React.StrictMode>
    <BrowserRouter>
      <NavigationScroll>
        <Toaster position="bottom-right" reverseOrder={false} />
        <Suspense fallback={<LazyLoad />}>
          <Pages />
        </Suspense>
      </NavigationScroll>
    </BrowserRouter>
  </React.StrictMode>,
  document.getElementById('root')
)

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals()
