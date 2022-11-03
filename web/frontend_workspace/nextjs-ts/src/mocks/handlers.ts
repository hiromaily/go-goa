import { rest } from 'msw'

export const handlers = [
  rest.get('/health', (req, res, ctx) => {
    return res(ctx.status(200), ctx.json({}))
  }),
  rest.get('/network', (req, res, ctx) => {
    return res(
      ctx.status(200),
      ctx.json([
        {
          id: 10,
          name: 'Optimism',
          rpcUrls: [
            'https://mainnet.optimism.io',
            'https://rpc.ankr.com/optimism',
            'https://optimism-mainnet.public.blastapi.io',
          ],
          blockExplorerUrls: ['https://optimistic.etherscan.io/'],
        },
        {
          id: 56,
          name: 'Binance Smart Chain Mainnet',
          rpcUrls: [
            'https://bsc-dataseed1.binance.org',
            'https://bsc-dataseed2.binance.org',
            'https://bsc-dataseed3.binance.org',
          ],
          blockExplorerUrls: ['https://bscscan.com/'],
        },
        {
          id: 137,
          name: 'Polygon Mainnet',
          rpcUrls: [
            'https://polygon-rpc.com',
            'https://rpc-mainnet.maticvigil.com',
            'https://rpc-mainnet.matic.network',
          ],
          blockExplorerUrls: ['https://polygonscan.com/'],
        },
      ]),
    )
  }),
  rest.get('/fee', (req, res, ctx) => {
    return res(
      ctx.status(200),
      ctx.json({
        gas: 0.05,
        bridgeFee: 1.5,
      }),
    )
  }),
]
