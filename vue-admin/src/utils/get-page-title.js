import defaultSettings from '@/settings'

const title = defaultSettings.title || 'Somewhere'

export default function getPageTitle(pageTitle) {
  if (pageTitle) {
    return `${pageTitle} - ${title}`
  }
  return `${title}`
}
