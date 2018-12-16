'use strict';

/**
 * Company.js controller
 *
 * @description: A set of functions called "actions" for managing `Company`.
 */

module.exports = {

  /**
   * Retrieve company records.
   *
   * @return {Object|Array}
   */

  find: async (ctx) => {
    if (ctx.query._q) {
      return strapi.services.company.search(ctx.query);
    } else {
      return strapi.services.company.fetchAll(ctx.query);
    }
  },

  /**
   * Retrieve a company record.
   *
   * @return {Object}
   */

  findOne: async (ctx) => {
    return strapi.services.company.fetch(ctx.params);
  },

  /**
   * Count company records.
   *
   * @return {Number}
   */

  count: async (ctx) => {
    return strapi.services.company.count(ctx.query);
  },

  /**
   * Create a/an company record.
   *
   * @return {Object}
   */

  create: async (ctx) => {
    return strapi.services.company.add(ctx.request.body);
  },

  /**
   * Update a/an company record.
   *
   * @return {Object}
   */

  update: async (ctx, next) => {
    return strapi.services.company.edit(ctx.params, ctx.request.body) ;
  },

  /**
   * Destroy a/an company record.
   *
   * @return {Object}
   */

  destroy: async (ctx, next) => {
    return strapi.services.company.remove(ctx.params);
  },

  /**
   * Add relation to a/an company record.
   *
   * @return {Object}
   */

  createRelation: async (ctx, next) => {
    return strapi.services.company.addRelation(ctx.params, ctx.request.body);
  },

  /**
   * Update relation to a/an company record.
   *
   * @return {Object}
   */

  updateRelation: async (ctx, next) => {
    return strapi.services.company.editRelation(ctx.params, ctx.request.body);
  },

  /**
   * Destroy relation to a/an company record.
   *
   * @return {Object}
   */

  destroyRelation: async (ctx, next) => {
    return strapi.services.company.removeRelation(ctx.params, ctx.request.body);
  }
};
