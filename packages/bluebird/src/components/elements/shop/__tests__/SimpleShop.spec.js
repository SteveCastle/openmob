import React from 'react';
import { shallow } from 'enzyme';
import SimpleShop from '../SimpleShop';

it('renders without crashing', () => {
  shallow(<SimpleShop />);
});
