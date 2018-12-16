'use strict';

/**
 * Promotion.js controller
 *
 * @description: A set of functions called "actions" for managing `Promotion`.
 */

module.exports = {

  /**
   * Retrieve promotion records.
   *
   * @return {Object|Array}
   */

  find: async (ctx) => {
    if (ctx.query._q) {
      return strapi.services.promotion.search(ctx.query);
    } else {
      return strapi.services.promotion.fetchAll(ctx.query);
    }
  },

  /**
   * Retrieve a promotion record.
   *
   * @return {Object}
   */

  findOne: async (ctx) => {
    return strapi.services.promotion.fetch(ctx.params);
  },

  /**
   * Count promotion records.
   *
   * @return {Number}
   */

  count: async (ctx) => {
    return strapi.services.promotion.count(ctx.query);
  },

  /**
   * Create a/an promotion record.
   *
   * @return {Object}
   */

  create: async (ctx) => {
    return strapi.services.promotion.add(ctx.request.body);
  },

  /**
   * Update a/an promotion record.
   *
   * @return {Object}
   */

  update: async (ctx, next) => {
    return strapi.services.promotion.edit(ctx.params, ctx.request.body) ;
  },

  /**
   * Destroy a/an promotion record.
   *
   * @return {Object}
   */

  destroy: async (ctx, next) => {
    return strapi.services.promotion.remove(ctx.params);
  },

  /**
   * Add relation to a/an promotion record.
   *
   * @return {Object}
   */

  createRelation: async (ctx, next) => {
    return strapi.services.promotion.addRelation(ctx.params, ctx.request.body);
  },

  /**
   * Update relation to a/an promotion record.
   *
   * @return {Object}
   */

  updateRelation: async (ctx, next) => {
    return strapi.services.promotion.editRelation(ctx.params, ctx.request.body);
  },

  /**
   * Destroy relation to a/an promotion record.
   *
   * @return {Object}
   */

  destroyRelation: async (ctx, next) => {
    return strapi.services.promotion.removeRelation(ctx.params, ctx.request.body);
  }
};
