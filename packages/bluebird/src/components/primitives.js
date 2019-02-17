import system from '@rebass/components';

export const Box = system(
  // core
  'space',
  'width',
  'color',
  'fontSize',
  // typography
  'textAlign',
  // borders
  'borders',
  'borderColor',
  'borderRadius',
  // layout
  'display',
  'maxWidth',
  'minWidth',
  'height',
  'maxHeight',
  'minHeight',
  // flexbox
  'alignItems',
  'alignContent',
  'justifyContent',
  'flexWrap',
  'flexDirection',
  'flex',
  'flexBasis',
  'justifySelf',
  'alignSelf',
  'order',
  // position
  'position',
  'zIndex',
  'top',
  'right',
  'bottom',
  'left'
);
Box.displayName = 'Box';

export const Text = system(
  {
    is: 'p',
    fontSize: 2,
    color: 'dark-gray',
    fontFamily: 'eczar',
    fontWeight: 400
  },
  // core
  'space',
  'width',
  'color',
  'fontSize',
  // typography
  'fontFamily',
  'textAlign',
  'lineHeight',
  'fontWeight',
  'letterSpacing'
);
Text.displayName = 'Text';

export const Heading = system(
  {
    is: 'h1',
    m: 0,
    fontSize: 6,
    color: 'dark-gray',
    fontFamily: 'eczar',
    fontWeight: 400
  },
  // core
  'space',
  'width',
  'color',
  'fontSize',
  // typography
  'fontFamily',
  'textAlign',
  'lineHeight',
  'fontWeight',
  'letterSpacing'
);
Heading.displayName = 'Heading';

export const BackgroundImage = system(
  {
    width: 1,
    ratio: 3 / 4,
    backgroundSize: 'cover',
    backgroundPosition: 'center',
    blacklist: ['src']
  },
  // core
  'space',
  'width',
  'color',
  'fontSize',
  'ratio',
  'backgroundSize',
  'backgroundPosition',
  props => ({
    backgroundImage: props.src ? `url(${props.src})` : undefined
  })
);
BackgroundImage.displayName = 'BackgroundImage';

export const Grid = system(
  // core
  'space',
  'width',
  'color',
  'fontSize',
  // typography
  'textAlign',
  // borders
  'borders',
  'borderColor',
  'borderRadius',
  // layout
  'maxWidth',
  'minWidth',
  'height',
  'maxHeight',
  'minHeight',
  // grid
  'gridGap',
  'gridColumnGap',
  'gridRowGap',
  'gridColumn',
  'gridRow',
  'gridAutoFlow',
  'gridAutoColumns',
  'gridAutoRows',
  'gridTemplateColumns',
  'gridTemplateRows',
  // flexbox
  'alignItems',
  'alignContent',
  'justifyContent',
  'justifySelf',
  {
    display: 'grid'
  }
);
Grid.displayName = 'Grid';

export const StrikeThrough = system(
  {
    is: 'span',
    blacklist: ['enabled']
  },
  'space',
  'width',
  'color',
  'fontSize',
  { position: 'relative' },
  props =>
    props.enabled && {
      '&:after': {
        content: '""',
        borderBottom: '1px solid currentcolor',
        left: 0,
        position: 'absolute',
        top: '35%',
        transform: 'skewY(11deg)',
        width: '100%'
      }
    }
);
