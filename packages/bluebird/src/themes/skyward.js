const colors = {
  base: '#3457a1',
  black: '#393f4a',
  white: '#fff',
  gray: [
    '#f9f9fa',
    '#ecedef',
    '#dee0e4',
    '#d0d2d8',
    '#c0c3ca',
    '#aeb3bc',
    '#9aa0ab',
    '#828997',
    '#636c7d',
    '#393f4a'
  ],
  blue: [
    '#e9edf5',
    '#d2daea',
    '#b7c4de',
    '#98aacf',
    '#7089bc',
    '#3458a1',
    '#2e4f90',
    '#28447d',
    '#213766',
    '#162646'
  ],
  indigo: [
    '#edebf5',
    '#d9d5eb',
    '#c2bce0',
    '#a79ed2',
    '#8377c0',
    '#4634a1',
    '#3e2e90',
    '#36287c',
    '#2b2064',
    '#1d1644'
  ],
  violet: [
    '#f2eaf5',
    '#e3d4eb',
    '#d3badf',
    '#bf9cd1',
    '#a674be',
    '#7d34a1',
    '#702e90',
    '#61287d',
    '#4f2166',
    '#371646'
  ],
  fuschia: [
    '#f5eaf3',
    '#ebd4e7',
    '#dfbad9',
    '#d09bc8',
    '#be73b2',
    '#a1348f',
    '#902e80',
    '#7e2870',
    '#67215c',
    '#481740'
  ],
  pink: [
    '#f5eaee',
    '#ebd4dc',
    '#dfbac6',
    '#d19cad',
    '#be748d',
    '#a13458',
    '#902e4f',
    '#7e2844',
    '#672138',
    '#471727'
  ],
  red: [
    '#f5ebea',
    '#ead6d2',
    '#debeb8',
    '#cfa198',
    '#bd7d70',
    '#a14634',
    '#903e2e',
    '#7e3628',
    '#672c21',
    '#481f17'
  ],
  orange: [
    '#f3efe6',
    '#e7decc',
    '#daccaf',
    '#cab68e',
    '#b89d66',
    '#a17d34',
    '#91702e',
    '#7e6228',
    '#685121',
    '#4a3917'
  ],
  yellow: [
    '#f0f3e5',
    '#e1e6c9',
    '#d0d8ab',
    '#bec889',
    '#a8b662',
    '#8fa134',
    '#81912e',
    '#717f29',
    '#5d6922',
    '#424b18'
  ],
  lime: [
    '#ebf3e6',
    '#d5e7cc',
    '#bcd9ae',
    '#a1ca8d',
    '#80b765',
    '#58a134',
    '#4f912e',
    '#457f29',
    '#396921',
    '#284a18'
  ],
  green: [
    '#e7f4e9',
    '#cde8d2',
    '#b0dab7',
    '#8fcb99',
    '#67b875',
    '#34a146',
    '#2e913f',
    '#297f37',
    '#21692d',
    '#184a20'
  ],
  teal: [
    '#e7f3ef',
    '#cce7de',
    '#afdacc',
    '#8ecab6',
    '#66b89d',
    '#34a17d',
    '#2e9170',
    '#297f62',
    '#216951',
    '#184b3a'
  ],
  cyan: [
    '#e7f2f4',
    '#cde3e8',
    '#b0d3da',
    '#8fc1cb',
    '#67abb9',
    '#348fa1',
    '#2e8191',
    '#29707e',
    '#215c68',
    '#18424a'
  ]
};

const theme = {
  breakpoints: ['40em', '52em', '64em'],
  space: [0, 4, 8, 16, 32, 64, 128, 256, 512],
  fontSizes: [12, 14, 16, 20, 24, 36, 48, 80, 96],
  fontWeights: [100, 200, 300, 400, 500, 600, 700, 800, 900],
  lineHeights: {
    solid: 1,
    title: 1.25,
    copy: 1.5
  },
  letterSpacings: {
    normal: 'normal',
    tracked: '0.1em',
    tight: '-0.05em',
    mega: '0.25em'
  },
  fonts: {
    serif: 'athelas, georgia, times, serif',
    sansSerif:
      '-apple-system, BlinkMacSystemFont, "avenir next", avenir, "helvetica neue", helvetica, ubuntu, roboto, noto, "segoe ui", arial, sans-serif'
  },
  borders: [
    0,
    '1px solid',
    '2px solid',
    '4px solid',
    '8px solid',
    '16px solid',
    '32px solid'
  ],
  radii: [0, 2, 4, 8, 16, 32, '100%'],
  widths: [16, 32, 64, 128, 256],
  heights: [16, 32, 64, 128, 256, '100%'],
  maxWidths: [16, 32, 64, 128, 256, 512, 768, 1024, 1536],
  shadows: ['0 0.5px 0 0 #ffffff inset, 0 1px 2px 0 #B3B3B3;'],
  colors: {
    ...colors,
    links: colors.blue[5],
    cellHover: colors.gray[1],
    modalBackground: 'rgba(0, 0, 0, .7)',

    linksHover: colors.blue[4],
    admin: {
      light: {
        bg: colors.gray[1],
        sidebarBg: colors.blue[9],
        menubarBg: colors.gray[1]
      },
      dark: { bg: colors.gray[8], sidebarBg: colors.blue[9] }
    },
    buttons: {
      light: {
        primary: {
          bg: colors.green[5],
          label: colors.white,
          hover: colors.green[4]
        }
      },
      dark: { primary: { bg: colors.gray[9], label: colors.white } }
    },
    forms: {
      light: {
        bg: colors.gray[2],
        value: colors.blue[9]
      },
      dark: { bg: colors.gray[9], value: colors.blue[9] }
    },
    type: {
      light: {
        header: colors.gray[9],
        subHeader: colors.gray[9],
        paragraph: colors.gray[9]
      },
      dark: {
        header: colors.gray[2],
        subHeader: colors.gray[2],
        paragraph: colors.gray[2]
      }
    }
  }
};

const skyward = {
  name: 'Skyward',
  ...theme
};

export default skyward;
