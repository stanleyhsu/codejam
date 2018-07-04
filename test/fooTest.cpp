#include "gtest/gtest.h"

#include "../foo.h"

using namespace std;

TEST(AddTest, OneAndTwo) {
    EXPECT_EQ(3, myadd(1, 2));
}

TEST(FooTest, AddOneAndTwo) {
    Foo *foo = new Foo();
    EXPECT_EQ(4, foo->add(2, 2));
}


TEST(FactTest, factN) {
    EXPECT_EQ(1, fact(0));
    EXPECT_EQ(1, fact(1));
    EXPECT_EQ(2, fact(2));
    EXPECT_EQ(6, fact(3));
}

TEST(FibTest, factN) {
    EXPECT_EQ(0, fib(0));
    EXPECT_EQ(1, fib(1));
    EXPECT_EQ(1, fib(2));
    EXPECT_EQ(2, fib(3));
    EXPECT_EQ(3, fib(4));
    EXPECT_EQ(5, fib(5));
    EXPECT_EQ(8, fib(6));
    EXPECT_EQ(13, fib(7));
//    EXPECT_EQ(12586269025, fib(50));
}

TEST(FibTest, factNwithMemo) {
    EXPECT_EQ(0, fibm(0));
    EXPECT_EQ(1, fibm(1));
    EXPECT_EQ(1, fibm(2));
    EXPECT_EQ(2, fibm(3));
    EXPECT_EQ(3, fibm(4));
    EXPECT_EQ(5, fibm(5));
    EXPECT_EQ(8, fibm(6));
    EXPECT_EQ(13, fibm(7));
//    EXPECT_EQ(12586269025, fibm(50));
}

int main(int argc, char **argv) {
    ::testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}