### 二叉树递归解决问题套路
1.  假设以X节点为头，假设可以向X左子树和右子树要任何信息
2.  在上一步的假设下，讨论以X为头节点的树，得到答案的可能性（最重要）
3.  列出所有的可能性后，确定到底需要向左子树和右子树要什么信息
4.  把左树的信息和右树的信息求全集，就是任何一棵子树都要返回的信息S
5.  递归函数都返回S，每一个子树都这么要求
6. 写代码，在代码中考虑如何把左树的信息和右树的信息整合出整棵树的信息