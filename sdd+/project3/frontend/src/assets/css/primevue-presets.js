export default {
  button: {
    root: ({ props }) => ({
      class: [
        'inline-flex items-center justify-center font-semibold text-sm transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-offset-2',
        // Sizing
        props.size === 'small' ? 'px-sm py-xs text-xs' : props.size === 'large' ? 'px-xl py-md text-base' : 'px-md py-sm',
        // Rounding
        'rounded-md',
        // Severity & Style Colors
        !props.severity || props.severity === 'primary'
           ? (props.link ? 'text-primary-600 hover:text-primary-700 bg-transparent' : props.outlined ? 'border border-primary-500 text-primary-500 hover:bg-primary-50 bg-transparent' : 'bg-primary-600 text-white hover:bg-primary-700')
          : props.severity === 'secondary'
          ? (props.outlined ? 'border border-secondary-500 text-secondary-500 hover:bg-slate-50 bg-transparent' : 'bg-secondary-500 text-white hover:bg-slate-600')
          : props.severity === 'success'
          ? (props.outlined ? 'border border-success-500 text-success-500 hover:bg-green-50 bg-transparent' : 'bg-success-500 text-white hover:bg-success-600')
          : props.severity === 'warning'
          ? (props.outlined ? 'border border-warning-500 text-warning-500 hover:bg-amber-50 bg-transparent' : 'bg-warning-500 text-white hover:bg-warning-600')
          : props.severity === 'danger'
          ? (props.outlined ? 'border border-danger-500 text-danger-500 hover:bg-red-50 bg-transparent' : 'bg-danger-500 text-white hover:bg-danger-600')
          : props.severity === 'info'
          ? (props.outlined ? 'border border-info-500 text-info-500 hover:bg-cyan-50 bg-transparent' : 'bg-info-500 text-white hover:bg-info-600')
          : '',
        // Disabled state
        props.disabled ? 'opacity-60 cursor-not-allowed' : 'cursor-pointer'
      ]
    })
  },
  inputtext: {
    root: ({ props }) => ({
      class: [
         'font-sans text-sm m-0 p-sm text-slate-700 bg-white border rounded-md transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500 w-full',
        props.invalid ? 'border-danger-500' : 'border-slate-300 hover:border-primary-500'
      ]
    })
  },
  select: {
    root: ({ props }) => ({
      class: [
         'inline-flex cursor-pointer select-none items-center justify-between border rounded-md transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-primary-500 bg-white p-sm text-sm w-full',
        props.invalid ? 'border-danger-500' : 'border-slate-300 hover:border-primary-500'
      ]
    }),
    label: { class: 'text-slate-700 font-sans' },
    trigger: { class: 'flex items-center justify-center text-slate-500 w-8' },
    panel: { class: 'bg-white border border-slate-200 rounded-md shadow-md mt-xs max-h-60 overflow-y-auto z-50' },
    list: { class: 'py-xs' },
    item: ({ context }) => ({
      class: [
        'cursor-pointer p-sm text-sm font-sans transition-colors duration-150',
        context.selected ? 'bg-primary-500 text-white font-semibold' : context.focused ? 'bg-slate-100 text-slate-700' : 'text-slate-700 hover:bg-slate-50'
      ]
    })
  },
  checkbox: {
    root: { class: 'inline-flex align-middle cursor-pointer select-none' },
    box: ({ props, context }) => ({
      class: [
        'flex items-center justify-center border w-5 h-5 rounded-md transition-colors duration-200',
        context.checked ? 'bg-primary-500 border-primary-500 text-white' : 'border-slate-300 bg-white',
        props.disabled ? 'opacity-60 cursor-not-allowed' : 'hover:border-primary-500'
      ]
    }),
    icon: { class: 'w-3 h-3' }
  },
  radiobutton: {
    root: { class: 'inline-flex align-middle cursor-pointer select-none' },
    box: ({ props, context }) => ({
      class: [
        'flex items-center justify-center border w-5 h-5 rounded-full transition-colors duration-200',
        context.checked ? 'bg-primary-500 border-primary-500 text-white' : 'border-slate-300 bg-white',
        props.disabled ? 'opacity-60 cursor-not-allowed' : 'hover:border-primary-500'
      ]
    }),
    icon: { class: 'w-2 h-2 rounded-full bg-white' }
  },
  card: {
    root: { class: 'bg-white rounded-lg border border-slate-200 shadow-soft overflow-hidden' },
    body: { class: 'p-lg flex flex-col gap-sm' },
    title: { class: 'text-lg font-bold text-slate-800' },
    subtitle: { class: 'text-xs text-slate-500 -mt-xs' },
    content: { class: 'text-sm text-slate-600' },
    footer: { class: 'flex justify-end gap-sm border-t border-slate-100 pt-md mt-sm' }
  },
  panel: {
    root: { class: 'bg-white border border-slate-200 rounded-lg overflow-hidden' },
    header: { class: 'flex items-center justify-between p-lg bg-slate-50 border-b border-slate-200 font-bold text-sm text-slate-700' },
    title: { class: 'text-sm font-bold' },
    content: { class: 'p-lg text-sm text-slate-600' }
  },
  dialog: {
    root: { class: 'bg-white rounded-lg border border-slate-200 shadow-hard overflow-hidden max-w-lg w-full flex flex-col z-50' },
    header: { class: 'flex items-center justify-between p-lg border-b border-slate-100' },
    title: { class: 'text-lg font-bold text-slate-800' },
    content: { class: 'p-lg text-sm text-slate-600 flex-grow' },
    footer: { class: 'flex justify-end gap-sm p-lg border-t border-slate-100' },
    mask: { class: 'fixed inset-0 z-40 bg-slate-900 bg-opacity-50 flex items-center justify-center p-md' }
  },
  datatable: {
    root: { class: 'relative overflow-x-auto border border-slate-200 rounded-lg bg-white' },
    table: { class: 'w-full border-collapse text-left text-sm text-slate-500' },
    thead: { class: 'bg-slate-50 text-xs font-semibold text-slate-700 uppercase border-b border-slate-200' },
    headerrow: { class: 'border-b border-slate-200' },
    column: {
      headercell: { class: 'px-lg py-md font-medium text-slate-700' },
      bodycell: { class: 'px-lg py-md border-b border-slate-100 text-slate-600' }
    },
    tbody: { class: 'divide-y divide-slate-100' },
    row: ({ context }) => ({
      class: [
        'transition-colors duration-150',
        context.selected ? 'bg-primary-50' : 'hover:bg-slate-50'
      ]
    })
  },
  tag: {
    root: ({ props }) => ({
      class: [
        'inline-flex items-center gap-xs px-2.5 py-0.5 rounded-full text-xs font-medium',
        !props.severity || props.severity === 'primary'
          ? 'bg-primary-50 text-primary-700 border border-primary-200'
          : props.severity === 'secondary'
          ? 'bg-slate-100 text-slate-700 border border-slate-200'
          : props.severity === 'success'
          ? 'bg-green-50 text-green-700 border border-green-200'
          : props.severity === 'warning'
          ? 'bg-amber-50 text-amber-700 border border-amber-200'
          : props.severity === 'danger'
          ? 'bg-red-50 text-red-700 border border-red-200'
          : props.severity === 'info'
          ? 'bg-cyan-50 text-cyan-700 border border-cyan-200'
          : ''
      ]
    })
  },
  badge: {
    root: ({ props }) => ({
      class: [
        'inline-flex items-center justify-center px-1.5 py-0.5 rounded-full text-xs font-semibold leading-none',
        !props.severity || props.severity === 'primary'
          ? 'bg-primary-500 text-white'
          : props.severity === 'secondary'
          ? 'bg-slate-500 text-white'
          : props.severity === 'success'
          ? 'bg-success-500 text-white'
          : props.severity === 'warning'
          ? 'bg-warning-500 text-white'
          : props.severity === 'danger'
          ? 'bg-danger-500 text-white'
          : props.severity === 'info'
          ? 'bg-info-500 text-white'
          : ''
      ]
    })
  },
  toast: {
    root: { class: 'fixed top-4 right-4 z-50 flex flex-col gap-sm w-96' },
    message: ({ props }) => ({
      class: [
        'p-lg rounded-lg border shadow-medium flex gap-md items-start transition-all duration-300',
        props.message.severity === 'success'
          ? 'bg-green-50 border-green-200 text-green-800'
          : props.message.severity === 'warn' || props.message.severity === 'warning'
          ? 'bg-amber-50 border-amber-200 text-amber-800'
          : props.message.severity === 'error' || props.message.severity === 'danger'
          ? 'bg-red-50 border-red-200 text-red-800'
          : 'bg-primary-50 border-primary-200 text-primary-800' // info/primary
      ]
    }),
    content: { class: 'flex-grow flex flex-col gap-xs' },
    summary: { class: 'text-sm font-bold' },
    detail: { class: 'text-xs opacity-90' },
    closebutton: { class: 'text-slate-400 hover:text-slate-600 transition-colors duration-150 p-xs rounded-md hover:bg-slate-100 -mt-xs -mr-xs ml-auto' }
  },
  menu: {
    root: { class: 'bg-white border border-slate-200 rounded-lg shadow-soft py-sm min-w-48' },
    list: { class: 'flex flex-col gap-xs px-sm' },
    item: { class: 'outline-none' },
    itemcontent: ({ context }) => ({
      class: [
        'flex items-center gap-sm px-md py-sm text-sm font-medium rounded-md transition-colors duration-150 cursor-pointer',
        context.focused ? 'bg-slate-100 text-slate-700' : 'text-slate-600 hover:bg-slate-50 hover:text-slate-800'
      ]
    })
  },
  breadcrumb: {
    root: { class: 'flex items-center py-sm bg-transparent' },
    list: { class: 'flex flex-wrap items-center gap-xs m-0 p-0 list-none text-sm text-slate-500' },
    item: { class: 'flex items-center gap-xs font-medium text-slate-600 hover:text-primary-600 cursor-pointer transition-colors duration-150' },
    separator: { class: 'flex items-center text-slate-400 mx-xs select-none' }
  }
};
