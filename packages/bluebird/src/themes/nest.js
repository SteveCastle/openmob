const colors = {
  base: '#43cea2',
  black: '#35413d',
  gray: [
    '#f8f9f9',
    '#e9eeed',
    '#dae2df',
    '#c9d5d1',
    '#b7c7c2',
    '#a3b7b1',
    '#8ca59d',
    '#738e86',
    '#5b7069',
    '#35413d'
  ],
  teal: [
    '#e7f8f3',
    '#cef2e6',
    '#b2ebd9',
    '#92e2c9',
    '#6ed9b7',
    '#43cea2',
    '#3cba92',
    '#35a380',
    '#2c876a',
    '#1f624d'
  ],
  cyan: [
    '#e7f5f8',
    '#ceebf2',
    '#b2e0eb',
    '#93d4e2',
    '#6ec6d9',
    '#43b5ce',
    '#3ca3ba',
    '#358fa3',
    '#2c7687',
    '#1f5561'
  ],
  blue: [
    '#eaeff9',
    '#d3ddf3',
    '#b9c9ec',
    '#9bb2e4',
    '#7696db',
    '#436fce',
    '#3c64b9',
    '#3457a1',
    '#2b4785',
    '#1e325d'
  ],
  indigo: [
    '#eeebfa',
    '#dcd6f4',
    '#c7beee',
    '#aea1e6',
    '#8e7cdd',
    '#5c43ce',
    '#523cb9',
    '#4734a1',
    '#3a2a83',
    '#281d5b'
  ],
  violet: [
    '#f5ebf9',
    '#ead5f4',
    '#debced',
    '#cf9fe6',
    '#bd7adc',
    '#a243ce',
    '#923cb9',
    '#7f34a2',
    '#682b85',
    '#4a1e5e'
  ],
  fuschia: [
    '#f9ebf7',
    '#f4d5ee',
    '#edbde5',
    '#e69fd9',
    '#dc7aca',
    '#ce43b5',
    '#b93ca3',
    '#a2348e',
    '#862b76',
    '#5f1f54'
  ],
  pink: [
    '#f9ebf0',
    '#f4d6df',
    '#edbdcc',
    '#e6a0b6',
    '#dc7b9a',
    '#ce436f',
    '#b93c64',
    '#a23457',
    '#862b48',
    '#5f1e33'
  ],
  red: [
    '#f9edea',
    '#f3d9d3',
    '#ecc3b9',
    '#e5a99b',
    '#db8876',
    '#ce5c43',
    '#b9533c',
    '#a24834',
    '#863b2b',
    '#5f2a1e'
  ],
  orange: [
    '#f8f2e6',
    '#f1e5cc',
    '#ead7af',
    '#e2c890',
    '#d8b66c',
    '#cea243',
    '#ba923c',
    '#a38035',
    '#876a2c',
    '#614c1f'
  ],
  yellow: [
    '#f4f8e5',
    '#eaf1c9',
    '#dee9ac',
    '#d2e18c',
    '#c4d869',
    '#b5ce43',
    '#a3ba3c',
    '#8fa335',
    '#77872c',
    '#56621f'
  ],
  lime: [
    '#ecf8e7',
    '#d8f2cd',
    '#c3eab0',
    '#aae291',
    '#8fd96d',
    '#6fce43',
    '#64ba3c',
    '#58a335',
    '#49872c',
    '#34611f'
  ],
  green: [
    '#e8f9eb',
    '#cff2d5',
    '#b3ebbd',
    '#94e3a2',
    '#70d983',
    '#43ce5c',
    '#3cba53',
    '#35a349',
    '#2c873c',
    '#1f622b'
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
  radii: [0, 2, 4, 16, 9999, '100%'],
  width: [16, 32, 64, 128, 256],
  heights: [16, 32, 64, 128, 256],
  maxWidths: [16, 32, 64, 128, 256, 512, 768, 1024, 1536],
  colors: {
    ...colors,
    darkBackground: colors.gray[9],
    lightText: colors.lime[0],
    lightBackground: colors.gray[1],
    darkText: colors.gray[9],
    buttons: {
      darkBackground: colors.gray[9],
      lightText: colors.lime[0],
      lightBackground: colors.gray[1],
      darkText: colors.gray[9]
    },
    forms: {
      darkBackground: colors.blue[9],
      lightText: colors.lime[0],
      lightBackground: colors.gray[1],
      darkText: colors.gray[9]
    },
    type: {
      darkBackground: colors.blue[9],
      lightText: colors.lime[0],
      lightBackground: colors.gray[1],
      darkText: colors.gray[9]
    }
  }
};

const nest = {
  name: 'Nest',
  ...theme
};

export default nest;
