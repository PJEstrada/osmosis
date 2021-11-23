"use strict";(self.webpackChunkdocs=self.webpackChunkdocs||[]).push([[404],{2194:(e,t,i)=>{i.r(t),i.d(t,{data:()=>o});const o={key:"v-5182764d",path:"/osmosis-app/liquidity-bootstraping.html",title:"LBPs",lang:"en-US",frontmatter:{},excerpt:"",headers:[{level:2,title:"Liquidity Bootstrapping Pools",slug:"liquidity-bootstrapping-pools",children:[]}],filePathRelative:"osmosis-app/liquidity-bootstraping.md",git:{updatedTime:1637704984e3,contributors:[{name:"Adam Tucker",email:"adamleetucker@outlook.com",commits:1}]}}},3199:(e,t,i)=>{i.r(t),i.d(t,{default:()=>g});var o=i(6252),a=i(2487);const s=(0,o._)("h1",{id:"lbps",tabindex:"-1"},[(0,o._)("a",{class:"header-anchor",href:"#lbps","aria-hidden":"true"},"#"),(0,o.Uk)(" LBPs")],-1),r=(0,o._)("h2",{id:"liquidity-bootstrapping-pools",tabindex:"-1"},[(0,o._)("a",{class:"header-anchor",href:"#liquidity-bootstrapping-pools","aria-hidden":"true"},"#"),(0,o.Uk)(" Liquidity Bootstrapping Pools")],-1),l=(0,o._)("p",null,"Osmosis offers a convenient design for Liquidity Bootstrapping Pools (LBPs), a special type of AMM designed for token sales. New protocols can use Osmosis’ LBP feature to distribute tokens and achieve initial price discovery.",-1),n=(0,o._)("p",null,"LBPs differ from other liquidity pools in terms of the ratio of assets within the pool. In LBPs, the ratio adjusts over time. LBPs involve an initial ratio, a target ratio, and an interval of time over which the ratio adjusts. These weights are customizable before launch. One can create an LBP with an initial ratio of 90-10, with the goal of reaching 50-50 over a month. The ratio continues to gradually adjust until the weights are equal within the pool. Like traditional LPs, the prices of assets within the pool is based on the ratio at the time of trade.",-1),h=(0,o._)("p",null,"After the LBP period ends or the final ratio is reached, the pool converts into a traditional LP pool.",-1),d=(0,o._)("p",null,"LBPs facilitate price discovery by demonstrating the acceptable market price of an asset. Ideally, LBPs will have very few buyers at the time of launch. The price slowly declines until traders are willing to step in and buy the asset. Arbitrage maintains this price for the remainder of the LBP. The process is shown by the blue line below. The price declines as the ratios adjust. Buyers step in until the acceptable price is again reached.",-1),c=(0,o._)("figure",null,[(0,o._)("img",{src:a,alt:""})],-1),p=(0,o._)("p",null,"Choosing the correct parameters is very important to price discovery for an LBP. If the initial price is too low, the asset will get bought up as soon as the pool is launched. It is also possible that the targeted final price is too high, creating little demand for the asset. The green line above shows this scenario. A buyer places a large order, but afterwards the price continues to decline as no additional buyers are willing to bite.",-1),u=(0,o._)("p",null,"Osmosis aims to provide an intuitive, easy-to-use LBP design to give protocols the best chance of a successful sale. The LBP feature facilitates initial price discovery for tokens and allows protocols to fairly distribute tokens to project stakeholders.",-1),f={},g=(0,i(3744).Z)(f,[["render",function(e,t){return(0,o.wg)(),(0,o.iD)(o.HY,null,[s,r,l,n,h,d,c,p,u],64)}]])},2487:(e,t,i)=>{e.exports=i.p+"assets/img/lbp.d36bfc21.png"}}]);